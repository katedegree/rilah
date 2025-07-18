package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/usecase/internal"
)

type getUserGroupsUsecase struct {
	groupRepository repository.GroupRepository
}

func NewGetUserGroupsUsecase(
	groupRepository repository.GroupRepository,
) *getUserGroupsUsecase {
	return &getUserGroupsUsecase{
		groupRepository: groupRepository,
	}
}

func (u *getUserGroupsUsecase) Execute(userID uint32) ([]*entity.GroupEntity, *internal.UsecaseError) {
	groups, err := u.groupRepository.ListByUserID(userID)
	if err != nil {
		return nil, &internal.UsecaseError{
			Message: "グループの取得に失敗しました",
		}
	}

	return groups, nil
}
