package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkScheds struct {
	ID        uint               `json:"id" gorm:"primaryKey"`
	Name      string             `json:"name"`
	Details   []WorkSchedDetails `json:"details" gorm:"foreignKey:WorkSchedID"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt gorm.DeletedAt     `json:"deleted_at" gorm:"index"`
}

type WorkSchedDetails struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	WorkSchedID uint           `json:"work_sched_id"`
	Day         string         `json:"day"`
	WorkStart   string         `json:"work_start"`
	WorkEnd     string         `json:"work_end"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
