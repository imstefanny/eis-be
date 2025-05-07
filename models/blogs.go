package models

import "time"

type Blogs struct {
	ID         uint      `json:"id"`
	Active     bool      `json:"active"`
	Title	   string    `json:"title"`
	Content    string    `json:"content"`
	Thumbnail  string    `json:"thumbnail"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedBy  uint	  	 `json:"created_by"`
	UpdatedBy  uint	  	 `json:"updated_by"`
}
