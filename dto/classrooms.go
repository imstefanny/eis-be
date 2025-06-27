package dto

import "gorm.io/gorm"

type CreateClassroomsRequest struct {
	LevelID   uint           `json:"level_id"`
	Grade     string         `json:"grade"`
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
