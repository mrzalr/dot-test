package models

import "github.com/google/uuid"

type Like struct {
	UserID string    `json:"user_id"`
	PostID uuid.UUID `gorm:"type:uuid"`
	Post   Post
}
