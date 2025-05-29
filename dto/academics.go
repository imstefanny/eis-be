package dto

import (
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
	Students          []int  `json:"students" validate:"required"`
}

type CreateBatchAcademicsRequest struct {
	StartYear string `json:"start_year" validate:"required"`
	EndYear   string `json:"end_year" validate:"required"`
}

type GetAcademicsResponse struct {
	ID              uint           `json:"id"`
	DisplayName     string         `json:"display_name"`
	Classroom       string         `json:"classroom"`
	Major           string         `json:"major"`
	HomeroomTeacher string         `json:"homeroom_teacher"`
	Students        int            `json:"students"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

type GetAcademicDetailResponse struct {
	ID              uint                         `json:"id"`
	DisplayName     string                       `json:"display_name"`
	StartYear       string                       `json:"start_year"`
	EndYear         string                       `json:"end_year"`
	Classroom       string                       `json:"classroom"`
	Major           string                       `json:"major"`
	HomeroomTeacher string                       `json:"homeroom_teacher"`
	Students        []GetStudentResponse         `json:"students"`
	SubjScheds      []GetSubjectScheduleResponse `json:"subject_schedules"`
	ClassNotes      []GetClassNoteResponse       `json:"class_notes"`
	// Attendances     []GetAttendanceResponse      `json:"attendances"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type GetStudentResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	NIS      string `json:"nis"`
}

type GetClassNoteResponse struct {
	Date    time.Time                   `json:"date"`
	Entries []GetClassNoteEntryResponse `json:"entries"`
}
type GetClassNoteEntryResponse struct {
	ID         uint   `json:"id"`
	Subject    string `json:"subject"`
	Teacher    string `json:"teacher"`
	TeacherAct string `json:"teacher_act"`
	Materials  string `json:"materials"`
	Notes      string `json:"notes"`
}
