package dto

type CreateStudentsRequest struct {
	ApplicantID       uint   `json:"applicant_id"`
	CurrentAcademicID uint   `json:"current_academic_id"`
	UserID            uint   `json:"user_id"`
	Email             string `json:"email"`
	ProfilePic        string `json:"profile_pic"`
	FullName          string `json:"full_name" validate:"required"`
	IdentityNo        string `json:"identity_no"`
	NIS               string `json:"nis"`
	NISN              string `json:"nisn"`
	PlaceOfBirth      string `json:"place_of_birth" validate:"required"`
	DateOfBirth       string `json:"date_of_birth" validate:"required"`
	Address           string `json:"address" validate:"required"`
	Phone             string `json:"phone"`
	Religion          string `json:"religion" validate:"required"`
	ChildSequence     int    `json:"child_sequence" validate:"required"`
	NumberOfSiblings  int    `json:"number_of_siblings" validate:"required"`
	LivingWith        string `json:"living_with" validate:"required"`
	ChildStatus       string `json:"child_status" validate:"required"`
}
