package dto

type CreateWorkSchedsRequest struct {
	Name    string                          `json:"name" validate:"required"`
	Details []CreateWorkSchedDetailsRequest `json:"details" validate:"required"`
}

type CreateWorkSchedDetailsRequest struct {
	Day         string `json:"day" validate:"required"`
	WorkStart   string `json:"work_start" validate:"required"`
	WorkEnd     string `json:"work_end" validate:"required"`
}
