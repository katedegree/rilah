package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/usecase/internal"
)

type getGroupUsersUsecase struct {
	userRepository repository.UserRepository
}

func NewGetGroupUsersUsecase(
	userRepository repository.UserRepository,
) *getGroupUsersUsecase {
	return &getGroupUsersUsecase{
		userRepository: userRepository,
	}
}

func (u *getGroupUsersUsecase) Execute(authID, groupID uint32) ([]*entity.UserEntity, *internal.UsecaseError) {
	users, err := u.userRepository.ListByGroupID(authID, groupID)
	if err != nil {
		return nil, &internal.UsecaseError{
			Code:    500,
			Message: "ユーザーの取得に失敗しました",
		}
	}

	return users, nil
}
