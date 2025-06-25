package models

import (
	"time"

	"gorm.io/gorm"
)

type Curriculums struct {
	ID                 uint                 `json:"id" gorm:"primaryKey"`
	DisplayName        string               `json:"display_name" gorm:"index"`
	Name               string               `json:"name"`
	LevelID            uint                 `json:"level_id" gorm:"index"`
	Level              Levels               `json:"level" gorm:"foreignKey:LevelID;references:ID"`
	Grade              string               `json:"grade"`
	CurriculumSubjects []CurriculumSubjects `json:"curriculum_subjects" gorm:"foreignKey:CurriculumID;references:ID"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
	DeletedAt          gorm.DeletedAt       `json:"deleted_at" gorm:"index"`
}

type CurriculumSubjects struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CurriculumID uint           `json:"curriculum_id" gorm:"size:255;uniqueIndex:curriculum_subjects"`
	SubjectID    uint           `json:"subject_id" gorm:"size:255;uniqueIndex:curriculum_subjects"`
	Subject      Subjects       `json:"subject" gorm:"foreignKey:SubjectID;references:ID"`
	Competence   string         `json:"competence"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
