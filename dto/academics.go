package dto

type CreateAcademicsRequest struct {
	DisplayName       string `json:"display_name"`
	StartYear         string `json:"start_year" validate:"required"`
	EndYear           string `json:"end_year" validate:"required"`
	ClassroomID       uint   `json:"classroom_id" validate:"required"`
	Major             string `json:"major"`
	HomeroomTeacherID uint   `json:"homeroom_teacher_id" validate:"required"`
	Students          []int  `json:"students" validate:"required"`
}

type CreateBatchAcademicsRequest struct {
	StartYear string `json:"start_year" validate:"required"`
	EndYear   string `json:"end_year" validate:"required"`
}

type GetAcademicsResponse struct {
	ID              uint   `json:"id"`
	DisplayName     string `json:"display_name"`
	Classroom       string `json:"classroom"`
	Major           string `json:"major"`
	HomeroomTeacher string `json:"homeroom_teacher"`
	Students        int    `json:"students"`
}
