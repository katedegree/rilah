package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/domain/service"
	"back/usecase/internal"
	"io"
)

type updateUserUsecase struct {
	userRepository repository.UserRepository
	fileRepository repository.FileRepository
}

func NewUpdateUserUsecase(userRepository repository.UserRepository, fileRepository repository.FileRepository) *updateUserUsecase {
	return &updateUserUsecase{
		userRepository: userRepository,
		fileRepository: fileRepository,
	}
}

func (u *updateUserUsecase) Execute(
	authUser *entity.UserEntity,
	name *string,
	accountCode *string,
	password *string,
	file io.ReadSeeker,
	contentType *string,
) *internal.UsecaseError {
	userEntity := &entity.UserEntity{
		ID: authUser.ID,
	}

	if name != nil {
		userEntity.Name = *name
	}
	if accountCode != nil {
		authService := service.NewAuthService()
		if authService.IsAccountCodeDuplicate(*accountCode, &authUser.AccountCode, u.userRepository) {
			return &internal.UsecaseError{
				Message: "アカウントコードが重複しています",
			}
		}
		if !authService.IsValidAccountCode(*accountCode) {
			return &internal.UsecaseError{
				Message: "アカウントコードの形式が正しくありません",
			}
		}
		userEntity.AccountCode = *accountCode
	}
	if password != nil {
		userEntity.Password = *password
	}
	if file != nil && contentType != nil {
		url, err := u.fileRepository.Upload(file, *contentType)
		if err != nil {
			return &internal.UsecaseError{
				Message: "画像ファイルのアップロードに失敗しました",
			}
		}
		userEntity.ImageURL = url
	}

	err := u.userRepository.Update(userEntity)
	if err != nil {
		return &internal.UsecaseError{
			Message: "ユーザー情報の更新に失敗しました",
		}
	}

	return nil
}
