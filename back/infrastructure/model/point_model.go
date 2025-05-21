package model

import (
	"time"

	"gorm.io/gorm"
)

func (PointModel) TableName() string {
	return "points"
}

type PointModel struct {
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	UserID  uint32 `gorm:"primaryKey;not null" json:"user_id"`
	GroupID uint32 `gorm:"primaryKey;not null" json:"group_id"`
	Amount  int32  `gorm:"not null" json:"amount"`

	User  UserModel  `gorm:"foreignKey:UserID" json:"user"`
	Group GroupModel `gorm:"foreignKey:GroupID" json:"group"`
}
