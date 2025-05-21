package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/domain/service"
	"back/usecase/internal"
)

type loginUsecase struct {
	userRepository        repository.UserRepository
	accessTokenRepository repository.AccessTokenRepository
}

func NewLoginUsecase(userRepository repository.UserRepository, accessTokenRepository repository.AccessTokenRepository) *loginUsecase {
	return &loginUsecase{
		userRepository:        userRepository,
		accessTokenRepository: accessTokenRepository,
	}
}

type loginUsecaseResponse struct {
	AccessToken string `json:"access_token"`
}

func (u *loginUsecase) Execute(ue entity.UserEntity) (*loginUsecaseResponse, *internal.UsecaseError) {
	userEntity, err := u.userRepository.FindByAccountCode(ue.AccountCode)
	if err != nil {
		return nil, &internal.UsecaseError{
			Message: "ログイン処理中に問題が発生しました。時間をおいて再試行してください。",
			Code:    500,
		}
	}
	if userEntity == nil {
		return nil, &internal.UsecaseError{
			Message: "メールアドレスまたはパスワードが正しくありません。",
			Code:    400,
		}
	}

	err = service.NewAuthService().ValidatePassword(userEntity.Password, ue.Password)
	if err != nil {
		return nil, &internal.UsecaseError{
			Message: "メールアドレスまたはパスワードが正しくありません。",
			Code:    400,
		}
	}

	accessTokenEntity, err := u.accessTokenRepository.Create(userEntity.ID)
	if err != nil {
		return nil, &internal.UsecaseError{
			Message: "ログイン処理中に問題が発生しました。時間をおいて再試行してください。",
			Code:    500,
		}
	}

	return &loginUsecaseResponse{
		AccessToken: accessTokenEntity.Token,
	}, nil
}
