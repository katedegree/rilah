package directive

import (
	"back/domain/constant"
	"back/domain/service"
	"back/infrastructure"
	"back/infrastructure/repository"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

func AuthDirective(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	req, ok := ctx.Value(constant.HTTP_REQUEST_KEY).(*http.Request)
	if !ok || req == nil {
		return nil, errors.New("failed to retrieve HTTP request from context")
	}

	tokenString := req.Header.Get("Authorization")
	if tokenString == "" {
		return nil, errors.New("authorization header missing")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	tokenString = strings.TrimSpace(tokenString)

	orm := infrastructure.Gorm()
	authService := service.NewAuthService()

	if !authService.ValidateToken(tokenString, repository.NewAccessTokenRepository(orm)) {
		return nil, errors.New("invalid or expired token")
	}

	authUser, err := repository.NewUserRepository(orm).FindByToken(tokenString)
	if err != nil || authUser == nil {
		return nil, errors.New("authenticated but user not found")
	}

	ctx = context.WithValue(ctx, constant.AUTH_USER_KEY, authUser)
	return next(ctx)
}
