package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.75

import (
	"back/domain/entity"
	"back/infrastructure/repository"
	"back/infrastructure/request"
	"back/usecase"
	"context"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, accountCode string, password string) (*entity.AuthResponse, error) {
	req := request.NewLoginRequest(accountCode, password)
	msgs, ok := req.Validate(r.Validator)
	if !ok {
		return &entity.AuthResponse{
			AccessToken: "",
			Success:     false,
			Messages:    msgs,
		}, nil
	}

	orm := r.Resolver.Orm

	loginUsecase := usecase.NewLoginUsecase(
		repository.NewUserRepository(orm),
		repository.NewAccessTokenRepository(orm),
	)
	loginUsecaseResponse, err := loginUsecase.Execute(
		entity.UserEntity{
			AccountCode: accountCode,
			Password:    password,
		})
	if err != nil {
		return &entity.AuthResponse{
			AccessToken: "",
			Success:     false,
			Messages:    []string{err.Message},
		}, nil
	}

	return &entity.AuthResponse{
		AccessToken: loginUsecaseResponse.AccessToken,
		Success:     true,
		Messages:    []string{"ログイン成功"},
	}, nil
}
