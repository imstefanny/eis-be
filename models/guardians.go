package models

import (
	"time"

	"gorm.io/gorm"
)

type Guardians struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	ApplicantID      uint           `json:"applicant_id"`
	Applicant        Applicants     `json:"applicant" gorm:"foreignKey:ApplicantID"`
	StudentID        uint           `json:"student_id"`
	Student          Students       `json:"student" gorm:"foreignKey:StudentID"`
	Relation         string         `json:"relation"`
	Name             string         `json:"name"`
	Religion         string         `json:"religion"`
	Job              string         `json:"job"`
	Address          string         `json:"address"`
	Phone            string         `json:"phone"`
	PlaceOfBirth     string         `json:"place_of_birth"`
	DateOfBirth      time.Time      `json:"date_of_birth"`
	HighestEducation string         `json:"highest_education"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
