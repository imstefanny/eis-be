package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentAttendances struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	DisplayName string         `json:"display_name" gorm:"index"`
	AcademicID  uint           `json:"academic_id" gorm:"default:null"`
	Academic    Academics      `json:"academic" gorm:"foreignKey:AcademicID"`
	StudentID   uint           `json:"student_id" gorm:"default:null"`
	Student     Students       `json:"student" gorm:"foreignKey:StudentID"`
	Date        time.Time      `json:"date" gorm:"index"`
	Status      string         `json:"status" gorm:"default:'Present'"`
	Remarks     string         `json:"remarks" gorm:"default:''"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
