package post

import "github.com/dot-test/internal/models"

type Repository interface {
	CreatePost(post models.Post) error
}

type Usecase interface {
	CreatePost(post models.Post) error
}
