package dto

import (
	"eis-be/models"
	"time"

	"gorm.io/gorm"
)

type CreateAcademicsRequest struct {
	DisplayName       string `json:"display_name"`
	StartYear         string `json:"start_year" validate:"required"`
	EndYear           string `json:"end_year" validate:"required"`
	ClassroomID       uint   `json:"classroom_id" validate:"required"`
	Major             string `json:"major"`
	HomeroomTeacherID uint   `json:"homeroom_teacher_id" validate:"required"`
	Students          []int  `json:"students"`
}

type CreateBatchAcademicsRequest struct {
	StartYear string `json:"start_year" validate:"required"`
	EndYear   string `json:"end_year" validate:"required"`
}

type GetAcademicsResponse struct {
	ID              uint              `json:"id"`
	DisplayName     string            `json:"display_name"`
	Classroom       string            `json:"classroom"`
	LevelName       string            `json:"level_name"`
	Major           string            `json:"major"`
	HomeroomTeacher string            `json:"homeroom_teacher"`
	Students        []models.Students `json:"students"`
	Terms           []models.Terms    `json:"terms"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `json:"deleted_at"`
}

type GetAcademicDetailResponse struct {
	ID                uint                         `json:"id"`
	DisplayName       string                       `json:"display_name"`
	StartYear         string                       `json:"start_year"`
	EndYear           string                       `json:"end_year"`
	Classroom         string                       `json:"classroom"`
	LevelName         string                       `json:"level_name"`
	Major             string                       `json:"major"`
	HomeroomTeacherId uint                         `json:"homeroom_teacher_id"`
	HomeroomTeacher   string                       `json:"homeroom_teacher"`
	Terms             []GetTermResponse            `json:"terms"`
	Students          []GetStudentResponse         `json:"students"`
	SubjScheds        []GetSubjectScheduleResponse `json:"subject_schedules"`
	ClassNotes        []GetClassNoteResponse       `json:"class_notes"`
	CreatedAt         time.Time                    `json:"created_at"`
	UpdatedAt         time.Time                    `json:"updated_at"`
	DeletedAt         gorm.DeletedAt               `json:"deleted_at"`
}

type GetTermResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetStudentResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	NIS      string `json:"nis"`
	NISN     string `json:"nisn"`
}

type StudentGetAcademicsResponse struct {
	ID          uint                     `json:"id"`
	DisplayName string                   `json:"display_name"`
	StartYear   string                   `json:"start_year"`
	EndYear     string                   `json:"end_year"`
	Terms       []StudentGetTermResponse `json:"terms"`
}
type StudentGetTermResponse struct {
	ID          uint   `json:"id"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
}
