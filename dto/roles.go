package dto

import (
	"time"

	"gorm.io/gorm"
)

type CreateRolesRequest struct {
	Name        string `json:"name" validate:"required"`
	Permissions []int  `json:"permissions" validate:"required"`
}

type GetRolesResponse struct {
	ID          int                      `json:"id"`
	Name        string                   `json:"name"`
	Permissions []GetPermissionsResponse `json:"permissions"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
	DeletedAt   gorm.DeletedAt           `json:"deleted_at,omitempty"`
}
type GetPermissionsResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
