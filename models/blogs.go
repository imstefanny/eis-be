package models

import (
	"time"

	"gorm.io/gorm"
)

type Blogs struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title"`
	Content       string         `json:"content"`
	Thumbnail     string         `json:"thumbnail"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedBy     uint           `json:"created_by"`
	CreatedByName Users          `json:"created_by_name" gorm:"foreignKey:CreatedBy"`
	UpdatedBy     uint           `json:"updated_by" gorm:"default:null"`
	UpdatedByName Users          `json:"updated_by_name" gorm:"foreignKey:UpdatedBy"`
}
