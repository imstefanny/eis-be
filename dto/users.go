package dto

type RegisterUsersRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}

type LoginUsersRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUsersResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Token  string `json:"token"`
	RoleID uint   `json:"role_id"`
}
