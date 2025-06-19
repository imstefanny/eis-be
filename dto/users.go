package dto

type RegisterUsersRequest struct {
	ProfilePic string `json:"profile_pic"`
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
}

type LoginUsersRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUsersResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Token       string   `json:"token"`
	RoleID      uint     `json:"role_id"`
	RoleName    string   `json:"role_name"`
	Permissions []string `json:"permissions"`
	ProfilePic  string   `json:"profile_pic"`
}

type UpdateUsersRequest struct {
	ProfilePic string `json:"profile_pic" validate:"omitempty,url"`
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password"`
	RoleID     uint   `json:"role_id" validate:"required,gt=0"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" validate:"required"`
}
