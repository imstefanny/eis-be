package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	ProfilePic string         `json:"profile_pic"`
	Name       string         `json:"name"`
	Email      string         `json:"email" gorm:"uniqueIndex;size:100"`
	Password   string         `json:"password"`
	RoleID     uint           `json:"role_id"`
	Role       Roles          `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
