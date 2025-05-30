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
	SubjSchedID uint   `json:"subj_sched_id" validate:"required"`
	TeacherID   uint   `json:"teacher_id" validate:"required"`
	Materials   string `json:"materials" validate:"required"`
	Notes       string `json:"notes" validate:"required"`
}

type CreateBatchClassNotesRequest struct {
	Date string `json:"date" validate:"required"`
}

type GetClassNotesResponse struct {
	ID         uint                        `json:"id"`
	AcademicID uint                        `json:"academic_id"`
	Date       time.Time                   `json:"date"`
	Details    []GetClassNoteEntryResponse `json:"details"`
	CreatedAt  time.Time                   `json:"created_at"`
	UpdatedAt  time.Time                   `json:"updated_at"`
	DeletedAt  gorm.DeletedAt              `json:"deleted_at"`
}
