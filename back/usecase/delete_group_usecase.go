package usecase

import (
	"back/domain/repository"
	"back/usecase/internal"
)

type deleteGroupUsecase struct {
	groupRepository repository.GroupRepository
}

func NewDeleteGroupUsecase(groupRepository repository.GroupRepository) *deleteGroupUsecase {
	return &deleteGroupUsecase{groupRepository: groupRepository}
}

func (u *deleteGroupUsecase) Execute(groupID, userID uint32) *internal.UsecaseError {
	err := u.groupRepository.Delete(groupID, userID)
	if err != nil {
		return &internal.UsecaseError{
			Message: "グループの削除に失敗しました",
		}
	}

	return nil
}
