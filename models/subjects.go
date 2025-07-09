package models

import (
	"time"

	"gorm.io/gorm"
)

type Subjects struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	DisplayName       string         `json:"display_name" gorm:"size:255;unique"`
	Code              string         `json:"code" gorm:"not null"`
	Name              string         `json:"name"`
	IsExtracurricular bool           `json:"is_extracurricular" gorm:"default:false"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
