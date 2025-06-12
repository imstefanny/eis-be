package models

import (
	"time"

	"gorm.io/gorm"
)

type Academics struct {
	ID                uint               `json:"id" gorm:"primaryKey"`
	DisplayName       string             `json:"display_name"`
	StartYear         string             `json:"start_year"`
	EndYear           string             `json:"end_year"`
	ClassroomID       uint               `json:"classroom_id" gorm:"default:null"`
	Classroom         Classrooms         `json:"classroom" gorm:"foreignKey:ClassroomID"`
	Major             string             `json:"major" gorm:"default:'General'"`
	HomeroomTeacherID uint               `json:"homeroom_teacher_id" gorm:"default:null"`
	HomeroomTeacher   Teachers           `json:"homeroom_teacher" gorm:"foreignKey:HomeroomTeacherID"`
	Students          []Students         `json:"students" gorm:"many2many:academic_students;"`
	SubjScheds        []SubjectSchedules `json:"subj_schedules" gorm:"foreignKey:AcademicID"`
	ClassNotes        []ClassNotes       `json:"class_notes" gorm:"foreignKey:AcademicID"`
	Terms             []Terms            `json:"terms" gorm:"foreignKey:AcademicID"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	DeletedAt         gorm.DeletedAt     `json:"deleted_at" gorm:"index"`
}

type Terms struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name"`
	AcademicID uint           `json:"academic_id"`
	Academic   Academics      `json:"academic" gorm:"foreignKey:AcademicID"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
