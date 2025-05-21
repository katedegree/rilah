package model

import (
	"time"

	"gorm.io/gorm"
)

func (AccessTokenModel) TableName() string {
	return "access_tokens"
}

type AccessTokenModel struct {
	ID        uint32         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	UserID    uint32    `gorm:"not null" json:"user_id"`
	Token     string    `gorm:"not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`

	User UserModel `gorm:"foreignKey:UserID" json:"user"`
}
