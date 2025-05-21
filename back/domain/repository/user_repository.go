package repository

import (
	"back/domain/entity"
)

type UserRepository interface {
	Create(ue *entity.UserEntity) (*entity.UserEntity, error)
	FindByAccountCode(account_code string) (*entity.UserEntity, error)
	FindByToken(token string) (*entity.UserEntity, error)
}
