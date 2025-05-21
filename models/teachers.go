package models

import (
	"time"

	"gorm.io/gorm"
)

type Teachers struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	IdentityNo  string         `json:"identity_no" gorm:"uniqueIndex;size:50"`
	Name        string         `json:"name"`
	NUPTK       string         `json:"nuptk" gorm:"uniqueIndex;size:50"`
	Phone       string         `json:"phone"`
	Email       string         `json:"email"`
	Address     string         `json:"address"`
	JobTitle    string         `json:"job_title"`
	LevelID     uint           `json:"level_id"`
	Level       Levels         `json:"level" gorm:"foreignKey:LevelID"`
	WorkSchedID uint           `json:"work_sched_id"`
	WorkSched   WorkScheds     `json:"work_sched" gorm:"foreignKey:WorkSchedID"`
	UserID      uint           `json:"user_id"`
	User        Users          `json:"user" gorm:"foreignKey:UserID"`
	ProfilePic  string         `json:"profile_pic"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
