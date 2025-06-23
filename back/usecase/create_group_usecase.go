package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/usecase/internal"
)

type createGroupUsecase struct {
	groupRepository       repository.GroupRepository
	pointRepository       repository.PointRepository
	transactionRepository repository.TransactionRepository
}

func NewCreateGroupUsecase(
	groupRepository repository.GroupRepository,
	pointRepository repository.PointRepository,
	transactionRepository repository.TransactionRepository,
) *createGroupUsecase {
	return &createGroupUsecase{
		groupRepository:       groupRepository,
		pointRepository:       pointRepository,
		transactionRepository: transactionRepository,
	}
}

func (u *createGroupUsecase) Execute(ge entity.GroupEntity, userID uint32) (*entity.GroupEntity, *internal.UsecaseError) {
	err := u.transactionRepository.ExecuteWith(func() error {
		groupEntity, err := u.groupRepository.Create(&ge, userID)
		if err != nil {
			return err
		}

		pe := entity.PointEntity{
			UserID:  userID,
			GroupID: groupEntity.ID,
		}
		_, err = u.pointRepository.Create(&pe)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, &internal.UsecaseError{
			Message: "グループの登録に失敗しました。",
		}
	}

	return &ge, nil
}
