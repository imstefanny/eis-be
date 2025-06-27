package dto

type CreateSubjectsRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type GetSubjectsResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
