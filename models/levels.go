package models

import (
	"time"

	"gorm.io/gorm"
)

type Levels struct {
	ID        uint             `json:"id" gorm:"primaryKey"`
	Name      string           `json:"name"`
	Histories []LevelHistories `json:"histories" gorm:"foreignKey:LevelID;references:ID"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at" gorm:"index"`
}

type LevelHistories struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	LevelID       uint      `json:"level_id"`
	OpCertNum     string    `json:"op_cert_num"`
	Accreditation string    `json:"accreditation"`
	Curriculum    string    `json:"curriculum"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	PrincipleID   uint      `json:"principle_id"`
	OperatorID    uint      `json:"operator_id"`
	State         bool      `json:"state"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
