package models

import (
	"time"

	"gorm.io/gorm"
)

type Academics struct {
	ID                uint               `json:"id" gorm:"primaryKey"`
	DisplayName       string             `json:"display_name" gorm:"size:255;unique"`
	StartYear         string             `json:"start_year"`
	EndYear           string             `json:"end_year"`
	ClassroomID       uint               `json:"classroom_id" gorm:"default:null"`
	Classroom         Classrooms         `json:"classroom" gorm:"foreignKey:ClassroomID"`
	Major             string             `json:"major" gorm:"default:'General'"`
	HomeroomTeacherID uint               `json:"homeroom_teacher_id" gorm:"default:null"`
	HomeroomTeacher   Teachers           `json:"homeroom_teacher" gorm:"foreignKey:HomeroomTeacherID"`
	CurriculumID      uint               `json:"curriculum_id" gorm:"default:null"`
	Curriculum        Curriculums        `json:"curriculum" gorm:"foreignKey:CurriculumID"`
	Students          []Students         `json:"students" gorm:"many2many:academic_students;"`
	AcademicStudents  []AcademicStudents `json:"academic_students" gorm:"foreignKey:AcademicsID"`
	SubjScheds        []SubjectSchedules `json:"subj_schedules" gorm:"foreignKey:AcademicID"`
	ClassNotes        []ClassNotes       `json:"class_notes" gorm:"foreignKey:AcademicID"`
	Terms             []Terms            `json:"terms" gorm:"foreignKey:AcademicID"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	DeletedAt         gorm.DeletedAt     `json:"deleted_at" gorm:"index"`
}

type Terms struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name" gorm:"size:255;uniqueIndex:academic_term"`
	AcademicID      uint           `json:"academic_id" gorm:"size:255;uniqueIndex:academic_term"`
	Academic        Academics      `json:"academic" gorm:"foreignKey:AcademicID"`
	FirstStartDate  time.Time      `json:"first_start_date" gorm:"default:null"`
	FirstEndDate    time.Time      `json:"first_end_date" gorm:"default:null"`
	SecondStartDate time.Time      `json:"second_start_date" gorm:"default:null"`
	SecondEndDate   time.Time      `json:"second_end_date" gorm:"default:null"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type AcademicStudents struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	AcademicsID     uint           `json:"academics_id" gorm:"index"`
	Academic        Academics      `json:"academic" gorm:"foreignKey:AcademicsID"`
	StudentsID      uint           `json:"students_id" gorm:"index"`
	Student         Students       `json:"student" gorm:"foreignKey:StudentsID"`
	FirstTermNotes  string         `json:"first_term_notes" gorm:"default:''"`
	SecondTermNotes string         `json:"second_term_notes" gorm:"default:''"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
