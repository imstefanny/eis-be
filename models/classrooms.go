package models

import (
	"time"

	"gorm.io/gorm"
)

type Classrooms struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	DisplayName string         `json:"display_name"`
	LevelID     uint           `json:"level_id"`
	Level       Levels         `json:"level" gorm:"foreignKey:LevelID"`
	Grade       string         `json:"grade"`
	Name        string         `json:"name"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
