package models

import (
	"time"

	"gorm.io/gorm"
)

type Academics struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	DisplayName       string         `json:"display_name"`
	StartYear         string         `json:"start_year"`
	EndYear           string         `json:"end_year"`
	ClassroomID       uint           `json:"classroom_id"`
	Major             string         `json:"major"`
	HomeroomTeacherID uint           `json:"homeroom_teacher_id"`
	Students		  []Students      `json:"students" gorm:"many2many:academic_students;"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
