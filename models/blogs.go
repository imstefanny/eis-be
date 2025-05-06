package models

import "time"

type Blogs struct {
	ID         uint      `json:"id"`
	Active     bool      `json:"active"`
	ProfilePic string    `json:"profile_pic"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	RoleID     uint      `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
