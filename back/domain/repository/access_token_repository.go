package repository

import (
	"back/domain/entity"
)

type AccessTokenRepository interface {
	Create(userId uint32) (*entity.AccessTokenEntity, error)

	FindByToken(token string) (*entity.AccessTokenEntity, error)
}
