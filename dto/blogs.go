package dto

type CreateBlogsRequest struct {
	Active     bool      `json:"active"`
	ProfilePic string    `json:"profile_pic"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	RoleID     uint      `json:"role_id"`
}
