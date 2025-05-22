package dto

type CreateAcademicsRequest struct {
	DisplayName       string `json:"display_name"`
	StartYear         string `json:"start_year" validate:"required"`
	EndYear           string `json:"end_year" validate:"required"`
	ClassroomID       uint   `json:"classroom_id" validate:"required"`
	Major             string `json:"major"`
	HomeroomTeacherID uint   `json:"homeroom_teacher_id" validate:"required"`
	Students          []int `json:"students" validate:"required"`
}
