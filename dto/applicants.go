package dto

type CreateApplicantsRequest struct {
	ProfilePic        string `json:"profile_pic"`
	FullName          string `json:"full_name" validate:"required"`
	IdentityNo        string `json:"identity_no"`
	Nisn              string `json:"nisn"`
	PlaceOfBirth      string `json:"place_of_birth" validate:"required"`
	DateOfBirth       string `json:"date_of_birth" validate:"required"`
	Address           string `json:"address" validate:"required"`
	Phone             string `json:"phone"`
	Religion          string `json:"religion" validate:"required"`
	ChildSequence     int    `json:"child_sequence" validate:"required"`
	NumberOfSiblings  int    `json:"number_of_siblings" validate:"required"`
	LivingWith        string `json:"living_with" validate:"required"`
	ChildStatus       string `json:"child_status" validate:"required"`
	SchoolOrigin      string `json:"school_origin"`
	LevelID           uint   `json:"level_id"`
	RegistrationGrade string `json:"registration_grade" validate:"required"`
	RegistrationMajor string `json:"registration_major"`
	State             string `json:"state" validate:"required"`
}

type RejectApplicantsRequest struct {
	Reason string `json:"reason" validate:"required"`
}
