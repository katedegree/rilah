package repository

import (
	"back/domain/entity"
)

type GroupRepository interface {
	Create(ge *entity.GroupEntity, authID uint32) (*entity.GroupEntity, error)
	FindByUserID(userID uint32) ([]*entity.GroupEntity, error)
	Update(ge *entity.GroupEntity, authID uint32) (*entity.GroupEntity, error)
	Delete(groupID, authID uint32) (*entity.GroupEntity, error)
	LinkUser(groupID, userID, authID uint32) (*entity.GroupEntity, error)
}
