package dto

type CreateClassroomsRequest struct {
	DisplayName string `json:"display_name"`
	LevelID     uint   `json:"level_id"`
	Grade       string `json:"grade"`
	Name        string `json:"name"`
}
