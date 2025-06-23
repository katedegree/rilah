package repository

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/infrastructure/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	orm *gorm.DB
}

func NewUserRepository(orm *gorm.DB) repository.UserRepository {
	return &userRepository{orm: orm}
}

func (r *userRepository) FindByAccountCode(accountCode string) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := r.orm.Where("account_code = ?", accountCode).First(&userModel)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user: %w", result.Error)
	}

	userEntity := &entity.UserEntity{
		ID:          userModel.ID,
		AccountCode: userModel.AccountCode,
		Password:    userModel.Password,
	}

	return userEntity, nil
}

func (r *userRepository) Create(user *entity.UserEntity) (*entity.UserEntity, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	userModel := &model.UserModel{
		Name:        user.Name,
		AccountCode: user.AccountCode,
		Password:    string(hashedPassword),
	}

	if err := r.orm.Create(userModel).Error; err != nil {
		return nil, err
	}

	return &entity.UserEntity{
		ID:          userModel.ID,
		Name:        userModel.Name,
		AccountCode: userModel.AccountCode,
	}, nil
}

func (r *userRepository) Update(user *entity.UserEntity) (*entity.UserEntity, error) {
	updateData := map[string]interface{}{}

	if user.Name != "" {
		updateData["name"] = user.Name
	}
	if user.AccountCode != "" {
		updateData["account_code"] = user.AccountCode
	}
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}
		updateData["password"] = string(hashedPassword)
	}
	if user.ImageURL != "" {
		updateData["image_url"] = user.ImageURL
	}

	if err := r.orm.Model(&model.UserModel{}).Where("id = ?", user.ID).Updates(updateData).Error; err != nil {
		return nil, err
	}

	var um model.UserModel
	if err := r.orm.First(&um, user.ID).Error; err != nil {
		return nil, err
	}

	ue := &entity.UserEntity{
		ID:          um.ID,
		Name:        um.Name,
		AccountCode: um.AccountCode,
		Password:    um.Password,
		ImageURL:    um.ImageURL,
	}
	return ue, nil
}

func (r *userRepository) FindByToken(token string) (*entity.UserEntity, error) {
	var tokenModel model.AccessTokenModel
	result := r.orm.Preload("User").Where("token = ?", token).First(&tokenModel)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user by token: %w", result.Error)
	}

	userEntity := &entity.UserEntity{
		ID:          tokenModel.User.ID,
		Name:        tokenModel.User.Name,
		AccountCode: tokenModel.User.AccountCode,
	}

	return userEntity, nil
}

func (r *userRepository) ListByGroupID(authID, groupID uint32) ([]*entity.UserEntity, error) {
	var user model.UserModel
	result := r.orm.Preload("Groups", "id = ?", groupID).First(&user, authID)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find user by ID: %w", result.Error)
	}
	if len(user.Groups) == 0 {
		return nil, fmt.Errorf("user does not belong to the group")
	}

	var group model.GroupModel
	result = r.orm.Preload("Users").First(&group, groupID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find group by ID: %w", result.Error)
	}

	userEntities := make([]*entity.UserEntity, len(group.Users))
	for i, userModel := range group.Users {
		userEntities[i] = &entity.UserEntity{
			ID:          userModel.ID,
			Name:        userModel.Name,
			AccountCode: userModel.AccountCode,
		}
	}

	return userEntities, nil
}
