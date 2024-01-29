package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Base
	Title       string
	Content     string
	IsPublished bool
}

func NewPost() Post {
	return Post{
		Base: Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:       "",
		Content:     "",
		IsPublished: false,
	}
}
