package models

import (
	"time"

	"gorm.io/gorm"
)

type Students struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	ApplicantID       uint           `json:"applicant_id" gorm:"index; default:null"`
	Applicant         Applicants     `json:"applicant" gorm:"foreignKey:ApplicantID;references:ID"`
	CurrentAcademicID uint           `json:"current_academic_id" gorm:"index"`
	Academics				 	Academics      `json:"academics" gorm:"foreignKey:CurrentAcademicID;references:ID"`
	UserID            uint           `json:"user_id" gorm:"index;default:null"`
	User              Users          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Guardians         []Guardians    `json:"guardians" gorm:"foreignKey:StudentID"`
	Documents         []Documents    `json:"documents" gorm:"foreignKey:StudentID"`
	ProfilePic        string         `json:"profile_pic"`
	FullName          string         `json:"full_name"`
	IdentityNo        string         `json:"identity_no"`
	NIS               string         `json:"nis"`
	NISN              string         `json:"nisn"`
	PlaceOfBirth      string         `json:"place_of_birth"`
	DateOfBirth       time.Time      `json:"date_of_birth"`
	Address           string         `json:"address"`
	Email             string         `json:"email"`
	Phone             string         `json:"phone"`
	Religion          string         `json:"religion"`
	ChildSequence     int            `json:"child_sequence"`
	NumberOfSiblings  int            `json:"number_of_siblings"`
	LivingWith        string         `json:"living_with"`
	ChildStatus       string         `json:"child_status"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
