package directive

import (
	"back/domain/entity"
	"back/domain/service"
	"back/infrastructure"
	"back/infrastructure/repository"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
)

type authDirective struct {
	Orm                *gorm.DB
	AuthUserContext    infrastructure.IContext[*entity.UserEntity]
	HttpRequestContext infrastructure.IContext[*http.Request]
}

func NewAuthDirective(orm *gorm.DB, authUesrContext infrastructure.IContext[*entity.UserEntity], HttpRequestContext infrastructure.IContext[*http.Request]) *authDirective {
	return &authDirective{
		Orm:                orm,
		AuthUserContext:    authUesrContext,
		HttpRequestContext: HttpRequestContext,
	}
}

func (d *authDirective) Execute(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	req := d.HttpRequestContext.Get(ctx)

	authHeader := req.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("AuthorizationヘッダーがBearerトークン形式ではありません")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	tokenString = strings.TrimSpace(tokenString)

	authService := service.NewAuthService()
	if !authService.ValidateToken(tokenString, repository.NewAccessTokenRepository(d.Orm)) {
		return nil, errors.New("トークンが無効または期限切れです")
	}

	authUser, err := repository.NewUserRepository(d.Orm).FindByToken(tokenString)
	if err != nil || authUser == nil {
		return nil, errors.New("該当するユーザーが見つかりませんでした")
	}

	ctx = d.AuthUserContext.Set(ctx, authUser)
	return next(ctx)
}
