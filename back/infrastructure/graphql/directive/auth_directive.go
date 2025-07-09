package directive

import (
	"back/domain/constant"
	"back/domain/service"
	"back/infrastructure/repository"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
)

type authDirective struct {
	orm *gorm.DB
}

func NewAuthDirective(orm *gorm.DB) *authDirective {
	return &authDirective{orm: orm}
}

func (d *authDirective) Execute(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	req, ok := ctx.Value(constant.HTTP_REQUEST_KEY).(*http.Request)
	if !ok || req == nil {
		return nil, errors.New("HTTPリクエスト情報の取得に失敗しました")
	}

	authHeader := req.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("AuthorizationヘッダーがBearerトークン形式ではありません")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	tokenString = strings.TrimSpace(tokenString)

	authService := service.NewAuthService()
	if !authService.ValidateToken(tokenString, repository.NewAccessTokenRepository(d.orm)) {
		return nil, errors.New("トークンが無効または期限切れです")
	}

	authUser, err := repository.NewUserRepository(d.orm).FindByToken(tokenString)
	if err != nil || authUser == nil {
		return nil, errors.New("該当するユーザーが見つかりませんでした")
	}

	ctx = context.WithValue(ctx, constant.AUTH_USER_KEY, authUser)
	return next(ctx)
}
