package models

import (
	"time"

	"gorm.io/gorm"
)

type Documents struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	TypeID       uint           `json:"type_id"`
	ApplicantID  uint           `json:"applicant_id"`
	StudentID    uint           `json:"student_id"`
	UploadedFile string         `json:"uploaded_file"`
	Description  string         `json:"description"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
