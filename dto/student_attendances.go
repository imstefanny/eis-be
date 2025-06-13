package dto

import (
	"time"

	"gorm.io/gorm"
)

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
	TermID     uint                            `json:"term_id"`
	Term       string                          `json:"term"`
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

type StudentGetAttendancesResponse struct {
	Month           int                       `json:"month"`
	Student         string                    `json:"student"`
	Academic        string                    `json:"academic"`
	PresenceCount   int                       `json:"presence_count"`
	PermissionCount int                       `json:"permission_count"`
	SickCount       int                       `json:"sick_count"`
	AlphaCount      int                       `json:"alpha_count"`
	Details         []StudentAttendanceDetail `json:"details"`
}
type StudentAttendanceDetail struct {
	ID      uint   `json:"id"`
	Date    string `json:"date"`
	Status  string `json:"status"`
	Remarks string `json:"remarks"`
}

type GetStudentAttsReport struct {
	Entries []GetStudentAttsDataReport  `json:"entries"`
	Levels  []GetStudentAttsLevelReport `json:"levels"`
}

type GetStudentAttsDataReport struct {
	Student         string `json:"student"`
	PresentCount    int    `json:"present_count"`
	PermissionCount int    `json:"permission_count"`
	SickCount       int    `json:"sick_count"`
	AlphaCount      int    `json:"alpha_count"`
}

type GetStudentAttsLevelReport struct {
	Level           string `json:"level"`
	PresentCount    int    `json:"present_count"`
	PermissionCount int    `json:"permission_count"`
	SickCount       int    `json:"sick_count"`
	AlphaCount      int    `json:"alpha_count"`
}
