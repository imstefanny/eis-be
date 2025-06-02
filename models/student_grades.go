package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentGrades struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	DisplayName string         `json:"display_name" gorm:"index"`
	AcademicID  uint           `json:"academic_id" gorm:"default:null"`
	Academic    Academics      `json:"academic" gorm:"foreignKey:AcademicID"`
	StudentID   uint           `json:"student_id" gorm:"default:null"`
	Student     Students       `json:"student" gorm:"foreignKey:StudentID"`
	SubjectID   uint           `json:"subject_id" gorm:"default:null"`
	Subject     Subjects       `json:"subject" gorm:"foreignKey:SubjectID"`
	Quiz        float64        `json:"quiz" gorm:"default:0"`
	FirstMonth  float64        `json:"first_month" gorm:"default:0"`
	SecondMonth float64        `json:"second_month" gorm:"default:0"`
	Finals      float64        `json:"finals" gorm:"default:0"`
	FinalGrade  float64        `json:"final_grade" gorm:"default:0"`
	Remarks     string         `json:"remarks" gorm:"default:''"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
