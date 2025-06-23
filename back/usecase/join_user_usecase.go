package usecase

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/domain/service"
	"back/usecase/internal"
)

type joinUserUsecase struct {
	groupRepository       repository.GroupRepository
	pointRepository       repository.PointRepository
	transactionRepository repository.TransactionRepository
}

func NewJoinUserUsecase(
	groupRepository repository.GroupRepository,
	pointRepository repository.PointRepository,
	transactionRepository repository.TransactionRepository,
) *joinUserUsecase {
	return &joinUserUsecase{
		groupRepository:       groupRepository,
		pointRepository:       pointRepository,
		transactionRepository: transactionRepository,
	}
}

func (u *joinUserUsecase) Execute(groupID, userID, authID uint32) (*entity.GroupEntity, *internal.UsecaseError) {
	var group *entity.GroupEntity

	err := u.transactionRepository.ExecuteWith(func() error {
		var err error
		group, err = u.groupRepository.LinkUser(groupID, userID, authID)
		if err != nil {
			return err
		}

		pe := &entity.PointEntity{
			UserID:  userID,
			GroupID: groupID,
		}

		err = service.NewPointService(u.pointRepository).EnsurePoint(pe)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, &internal.UsecaseError{
			Message: "ユーザーの招待に失敗しました",
		}
	}

	return group, nil
}
