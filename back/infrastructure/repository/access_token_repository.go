package repository

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/infrastructure/model"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type accessTokenRepository struct {
	orm *gorm.DB
}

func NewAccessTokenRepository(orm *gorm.DB) repository.AccessTokenRepository {
	return &accessTokenRepository{orm: orm}
}

func (r *accessTokenRepository) Create(userId uint32) (*entity.AccessTokenEntity, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"jti":     uuid.New().String(),
	})

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	accessToken := &model.AccessTokenModel{
		UserID:    userId,
		Token:     signedToken,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := r.orm.Create(accessToken).Error; err != nil {
		return nil, err
	}

	return &entity.AccessTokenEntity{
		ID:        accessToken.ID,
		UserID:    accessToken.UserID,
		Token:     accessToken.Token,
		ExpiresAt: accessToken.ExpiresAt,
	}, nil
}

func (r *accessTokenRepository) FindByToken(token string) (*entity.AccessTokenEntity, error) {
	var accessTokenModel model.AccessTokenModel
	result := r.orm.Where("token = ?", token).First(&accessTokenModel)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find access token: %w", result.Error)
	}

	accessTokenEntity := &entity.AccessTokenEntity{
		ID:        accessTokenModel.ID,
		UserID:    accessTokenModel.UserID,
		Token:     accessTokenModel.Token,
		ExpiresAt: accessTokenModel.ExpiresAt,
	}

	return accessTokenEntity, nil
}
