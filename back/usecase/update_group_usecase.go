package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/usecase/internal"
)

type updateGroupUsecase struct {
	groupRepository repository.GroupRepository
}

func NewUpdateGroupUsecase(groupRepository repository.GroupRepository) *updateGroupUsecase {
	return &updateGroupUsecase{groupRepository: groupRepository}
}

func (u *updateGroupUsecase) Execute(ge *entity.GroupEntity, userID uint32) *internal.UsecaseError {
	err := u.groupRepository.Update(ge, userID)
	if err != nil {
		return &internal.UsecaseError{
			Message: "グループの更新に失敗しました",
		}
	}

	return nil
}
