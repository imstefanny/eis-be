package dto

type CreateGuardiansRequest struct {
	ApplicantID      uint   `json:"applicant_id"`
	StudentID        uint   `json:"student_id"`
	Relation         string `json:"relation" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Religion         string `json:"religion" validate:"required"`
	Job              string `json:"job" validate:"required"`
	Address          string `json:"address" validate:"required"`
	Phone            string `json:"phone" validate:"required"`
	PlaceOfBirth     string `json:"place_of_birth" validate:"required"`
	DateOfBirth      string `json:"date_of_birth" validate:"required"`
	HighestEducation string `json:"highest_education" validate:"required"`
}

// {
//     "applicant_id": 1,
//     "student_id": 1,
//     "relation": "daddy",
//     "name": "Budi",
//     "religion": "Christian",
//     "job": "Programmer",
//     "address": "Jalan jalan No 1",
//     "phone": "081274584134",
//     "place_of_birth": "Medan",
//     "date_of_birth": "1995-05-10",
//     "highest_education": "S1"
// }
