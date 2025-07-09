package dto

type CreateSubjectsRequest struct {
	Code              string `json:"code"`
	Name              string `json:"name"`
	IsExtracurricular bool   `json:"is_extracurricular"`
}

type GetSubjectsResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Code              string `json:"code"`
	IsExtracurricular bool   `json:"is_extracurricular"`
}
