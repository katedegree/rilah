package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/domain/service"
	"back/usecase/internal"
)

type signUpUsecase struct {
	userRepository        repository.UserRepository
	accessTokenRepository repository.AccessTokenRepository
}

func NewSignUpUsecase(userRepository repository.UserRepository, accessTokenRepository repository.AccessTokenRepository) *signUpUsecase {
	return &signUpUsecase{
		userRepository:        userRepository,
		accessTokenRepository: accessTokenRepository,
	}
}

type signUpUsecaseResponse struct {
	AccessToken string `json:"access_token"`
}

func (u *signUpUsecase) Execute(ue entity.UserEntity) (*signUpUsecaseResponse, *internal.UsecaseError) {
	authService := service.NewAuthService()

	if authService.IsAccountCodeDuplicate(ue.AccountCode, u.userRepository) {
		return nil, &internal.UsecaseError{
			Code:    400,
			Message: "既に登録されているアカウントコードです。",
		}
	}

	if !authService.IsValidAccountCode(ue.AccountCode) {
		return nil, &internal.UsecaseError{
			Code:    400,
			Message: "アカウントコードはa-z, A-Z, 0-9, -, _のみ使用できます。",
		}
	}

	user, err := u.userRepository.Create(&ue)
	if err != nil {
		return nil, &internal.UsecaseError{
			Code:    500,
			Message: "ユーザー作成に失敗しました。",
		}
	}

	accessToken, err := u.accessTokenRepository.Create(user.ID)
	if err != nil {
		return nil, &internal.UsecaseError{
			Code:    500,
			Message: "ログインに失敗しました。再度ログインしてください。",
		}
	}

	return &signUpUsecaseResponse{
		AccessToken: accessToken.Token,
	}, nil
}
