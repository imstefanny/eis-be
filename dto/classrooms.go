package dto

type CreateClassroomsRequest struct {
	LevelID     uint   `json:"level_id"`
	Grade       string `json:"grade"`
	Name        string `json:"name"`
}
