package repository

import (
	"back/domain/entity"
	"back/domain/repository"
	"back/infrastructure/model"
	"errors"

	"gorm.io/gorm"
)

type groupRepository struct {
	orm *gorm.DB
}

func NewGroupRepository(orm *gorm.DB) repository.GroupRepository {
	return &groupRepository{orm: orm}
}

func (r *groupRepository) Create(ge *entity.GroupEntity, authID uint32) (*entity.GroupEntity, error) {
	var userModel model.UserModel

	if err := r.orm.First(&userModel, authID).Error; err != nil {
		return nil, err
	}

	groupModel := model.GroupModel{
		Name: ge.Name,
	}

	if err := r.orm.Create(&groupModel).Error; err != nil {
		return nil, err
	}

	if err := r.orm.Model(&userModel).Association("Groups").Append(&groupModel); err != nil {
		return nil, err
	}

	return &entity.GroupEntity{
		ID:   groupModel.ID,
		Name: groupModel.Name,
	}, nil
}

func (r *groupRepository) ListByUserID(userID uint32) ([]*entity.GroupEntity, error) {
	var user model.UserModel

	if err := r.orm.Preload("Groups").First(&user, userID).Error; err != nil {
		return nil, err
	}

	groupEntities := make([]*entity.GroupEntity, 0, len(user.Groups))
	for _, gm := range user.Groups {
		groupEntities = append(groupEntities, &entity.GroupEntity{
			ID:        gm.ID,
			Name:      gm.Name,
			CreatedAt: gm.CreatedAt,
			UpdatedAt: gm.UpdatedAt,
		})
	}

	return groupEntities, nil
}

func (r *groupRepository) Update(ge *entity.GroupEntity, authID uint32) (*entity.GroupEntity, error) {
	var user model.UserModel
	if err := r.orm.Preload("Groups").First(&user, authID).Error; err != nil {
		return nil, err
	}

	var groupModel *model.GroupModel
	for _, group := range user.Groups {
		if group.ID == ge.ID {
			groupModel = &group
			break
		}
	}

	if groupModel == nil {
		return nil, gorm.ErrRecordNotFound
	}

	groupModel.Name = ge.Name
	if err := r.orm.Save(groupModel).Error; err != nil {
		return nil, err
	}

	return &entity.GroupEntity{
		ID:        groupModel.ID,
		Name:      groupModel.Name,
		CreatedAt: groupModel.CreatedAt,
		UpdatedAt: groupModel.UpdatedAt,
	}, nil
}

func (r *groupRepository) Delete(groupID, authID uint32) (*entity.GroupEntity, error) {
	var user model.UserModel
	if err := r.orm.Preload("Groups").First(&user, authID).Error; err != nil {
		return nil, err
	}

	var groupToDelete *model.GroupModel
	for _, group := range user.Groups {
		if group.ID == groupID {
			groupToDelete = &group
			break
		}
	}

	if groupToDelete == nil {
		return nil, gorm.ErrRecordNotFound
	}

	deletedEntity := &entity.GroupEntity{
		ID:        groupToDelete.ID,
		Name:      groupToDelete.Name,
		CreatedAt: groupToDelete.CreatedAt,
		UpdatedAt: groupToDelete.UpdatedAt,
	}

	if err := r.orm.Delete(&model.GroupModel{}, groupID).Error; err != nil {
		return nil, err
	}

	return deletedEntity, nil
}

func (r *groupRepository) LinkUser(groupID, userID, authID uint32) (*entity.GroupEntity, error) {
	var authUser model.UserModel
	if err := r.orm.Preload("Groups", "id = ?", groupID).First(&authUser, authID).Error; err != nil {
		return nil, err
	}
	if len(authUser.Groups) == 0 {
		return nil, errors.New("指定された groupID は認証ユーザーに紐づいていません")
	}
	groupModel := authUser.Groups[0]

	var invitedUser model.UserModel
	if err := r.orm.First(&invitedUser, userID).Error; err != nil {
		return nil, err
	}

	if err := r.orm.Model(&invitedUser).Association("Groups").Append(&groupModel); err != nil {
		return nil, err
	}

	return &entity.GroupEntity{
		ID:        groupModel.ID,
		Name:      groupModel.Name,
		CreatedAt: groupModel.CreatedAt,
		UpdatedAt: groupModel.UpdatedAt,
	}, nil
}
