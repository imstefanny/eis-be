package models

import (
	"time"

	"gorm.io/gorm"
)

type SubjectSchedules struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	DisplayName string              `json:"display_name"`
	AcademicID  uint                `json:"academic_id"`
	Academic    Academics           `json:"academic" gorm:"foreignKey:AcademicID"`
	SubjectID   uint                `json:"subject_id"`
	Subject     Subjects            `json:"subject" gorm:"foreignKey:SubjectID"`
	TeacherID   uint                `json:"teacher_id" gorm:"default:null"`
	Teacher     Teachers            `json:"teacher" gorm:"foreignKey:TeacherID"`
	Day         string              `json:"day"`
	StartHour   string              `json:"start_hour"`
	EndHour     string              `json:"end_hour"`
	Histories   []ClassNotesDetails `json:"histories" gorm:"foreignKey:SubjSchedID;references:ID"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
}
