package dto

import (
	"time"

	"gorm.io/gorm"
)

type CreateClassNotesRequest struct {
	AcademicID uint                             `json:"academic_id" validate:"required"`
	Date       string                           `json:"date" validate:"required"`
	Details    []CreateClassNotesDetailsRequest `json:"details" validate:"required"`
}

type CreateClassNotesDetailsRequest struct {
	ID          uint   `json:"id"`
	NoteID      uint   `json:"note_id"`
	SubjSchedID uint   `json:"subj_sched_id" validate:"required"`
	TeacherID   uint   `json:"teacher_id" validate:"required"`
	Materials   string `json:"materials" validate:"required"`
	Notes       string `json:"notes" validate:"required"`
}

type CreateBatchClassNotesRequest struct {
	Date string `json:"date" validate:"required"`
}

type GetClassNoteResponse struct {
	ID      uint                        `json:"id"`
	Date    time.Time                   `json:"date"`
	Entries []GetClassNoteEntryResponse `json:"entries"`
}
type GetClassNoteEntryResponse struct {
	ID                uint   `json:"id"`
	Subject           string `json:"subject"`
	SubjectScheduleId uint   `json:"subject_schedule_id"`
	Teacher           string `json:"teacher"`
	TeacherID         uint   `json:"teacher_id"`
	TeacherAct        string `json:"teacher_act"`
	TeacherActID      uint   `json:"teacher_act_id"`
	Materials         string `json:"materials"`
	Notes             string `json:"notes"`
}

type GetClassNotesResponse struct {
	ID             uint                          `json:"id"`
	AcademicID     uint                          `json:"academic_id"`
	Date           time.Time                     `json:"date"`
	Details        []GetClassNoteEntryResponse   `json:"details"`
	AbsenceCount   []GetClassNoteAbsenceResponse `json:"absence_count"`
	AbsenceDetails []GetClassNoteAbsenceDetails  `json:"absence_details"`
	CreatedAt      time.Time                     `json:"created_at"`
	UpdatedAt      time.Time                     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt                `json:"deleted_at"`
}
type GetClassNoteAbsenceResponse struct {
	Status string `json:"status"`
	Total  int    `json:"total"`
}

type GetClassNoteAbsenceDetails struct {
	ID        uint   `json:"attendance_id"`
	StudentID uint   `json:"student_id"`
	FullName  string `json:"full_name"`
	Status    string `json:"status"`
	Remarks   string `json:"remarks"`
}

type BrowseClassNotesResponse struct {
	ID          uint           `json:"id"`
	DisplayName string         `json:"display_name"`
	AcademicID  uint           `json:"academic_id"`
	Academic    string         `json:"academic"`
	Date        time.Time      `json:"date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
