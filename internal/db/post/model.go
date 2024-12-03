package post

import (
	"time"
)

type ModelPost struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Category  string     `json:"category"`
	Tags      []string   `json:"tags"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
