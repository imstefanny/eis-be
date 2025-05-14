package dto

type CreateWorkSchedsRequest struct {
	Name    string                          `json:"name" validate:"required"`
	Details []CreateWorkSchedDetailsRequest `json:"details" validate:"required"`
}

type CreateWorkSchedDetailsRequest struct {
	ID	uint	`json:"id"`
	Day         string `json:"day" validate:"required"`
	WorkStart   string `json:"work_start" validate:"required"`
	WorkEnd     string `json:"work_end" validate:"required"`
}
