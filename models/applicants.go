package models

import (
	"time"

	"gorm.io/gorm"
)

type Applicants struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	ProfilePic        string         `json:"profile_pic"`
	FullName          string         `json:"full_name"`
	IdentityNo        string         `json:"identity_no"`
	PlaceOfBirth      string         `json:"place_of_birth"`
	DateOfBirth       time.Time      `json:"date_of_birth"`
	Address           string         `json:"address"`
	Phone             string         `json:"phone"`
	Religion          string         `json:"religion"`
	ChildSequence     int            `json:"child_sequence"`
	NumberOfSiblings  int            `json:"number_of_siblings"`
	LivingWith        string         `json:"living_with"`
	ChildStatus       string         `json:"child_status"`
	SchoolOrigin      string         `json:"school_origin"`
	LevelID           uint           `json:"level_id"`
	RegistrationGrade string         `json:"registration_grade"`
	RegistrationMajor string         `json:"registration_major"`
	State             string         `json:"state"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedBy         uint           `json:"created_by"`
	UpdatedBy         uint           `json:"updated_by"`
}
