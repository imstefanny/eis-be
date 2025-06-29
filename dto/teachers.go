package dto

import (
	"gorm.io/gorm"
)

type CreateTeachersRequest struct {
	IdentityNo  string         `json:"identity_no" validate:"required"`
	Name        string         `json:"name" validate:"required"`
	NUPTK       string         `json:"nuptk"`
	Phone       string         `json:"phone" validate:"required"`
	Email       string         `json:"email" validate:"required,email"`
	Address     string         `json:"address" validate:"required"`
	JobTitle    string         `json:"job_title" validate:"required"`
	LevelID     uint           `json:"level_id"`
	RoleID      uint           `json:"role_id" validate:"required"`
	WorkSchedID uint           `json:"work_sched_id" validate:"required"`
	ProfilePic  string         `json:"profile_pic"`
	MachineID   uint           `json:"machine_id" validate:"required"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"`
}
