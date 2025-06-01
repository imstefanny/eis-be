package dto

import (
	"time"

	"gorm.io/gorm"
)

// type CreateStudentAttsRequest struct {
// 	AcademicID uint   `json:"academic_id" validate:"required"`
// 	StudentID  uint   `json:"student_id" validate:"required"`
// 	Date       string `json:"date" validate:"required"`
// 	Status     string `json:"status" validate:"required"`
// 	Remarks    string `json:"remarks" validate:"required"`
// }

type CreateBatchStudentAttsRequest struct {
	Date string `json:"date" validate:"required"`
}

type UpdateStudentAttsRequest struct {
	AcademicID uint                            `json:"academic_id" validate:"required"`
	Date       string                          `json:"date" validate:"required"`
	Students   []UpdateStudentAttsEntryRequest `json:"students" validate:"required,dive"`
}
type UpdateStudentAttsEntryRequest struct {
	StudentID uint   `json:"student_id" validate:"required"`
	Status    string `json:"status" validate:"required"`
	Remarks   string `json:"remarks" validate:"required"`
}

type GetAllStudentAttsRequest struct {
	AcademicID uint                            `json:"academic_id"`
	Academic   string                          `json:"academic"`
	Date       string                          `json:"date"`
	Students   []GetAllStudentAttsEntryRequest `json:"students"`
}
type GetAllStudentAttsEntryRequest struct {
	ID          uint           `json:"id"`
	StudentID   uint           `json:"student_id"`
	Student     string         `json:"student"`
	DisplayName string         `json:"display_name"`
	Status      string         `json:"status"`
	Remarks     string         `json:"remarks"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
