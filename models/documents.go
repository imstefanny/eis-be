package models

import (
	"time"

	"gorm.io/gorm"
)

type Documents struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	TypeID       uint           `json:"type_id"`
	Type         DocTypes       `json:"type" gorm:"foreignKey:TypeID"`
	ApplicantID  uint           `json:"applicant_id" gorm:"default:null"`
	Applicant    Applicants     `json:"applicant" gorm:"foreignKey:ApplicantID"`
	StudentID    uint           `json:"student_id" gorm:"default:null"`
	Student      Students       `json:"student" gorm:"foreignKey:StudentID"`
	UploadedFile string         `json:"uploaded_file"`
	Description  string         `json:"description"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
