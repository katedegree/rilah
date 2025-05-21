package repository

import (
	"back/domain/entity"
)

type PointRepository interface {
	Create(pe *entity.PointEntity) (*entity.PointEntity, error)
	FindByUserAndGroup(pe *entity.PointEntity) (*entity.PointEntity, error)
	Restore(pe *entity.PointEntity) (*entity.PointEntity, error)
}
