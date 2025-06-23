package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/usecase/internal"
)

type deleteGroupUsecase struct {
	groupRepository repository.GroupRepository
}

func NewDeleteGroupUsecase(groupRepository repository.GroupRepository) *deleteGroupUsecase {
	return &deleteGroupUsecase{groupRepository: groupRepository}
}

func (u *deleteGroupUsecase) Execute(groupID, userID uint32) (*entity.GroupEntity, *internal.UsecaseError) {
	group, err := u.groupRepository.Delete(groupID, userID)
	if err != nil {
		return nil, &internal.UsecaseError{
			Message: "グループの削除に失敗しました",
		}
	}

	return group, nil
}
