package models

import (
	"time"

	"gorm.io/gorm"
)

type Subjects struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Code      string         `json:"code" gorm:"uniqueIndex;not null"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
