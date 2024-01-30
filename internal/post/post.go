package post

import (
	"github.com/dot-test/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	CreatePost(post models.Post) error
	GetPostByID(id uuid.UUID) (models.Post, error)
	GetPosts() ([]models.Post, error)
	UpdatePost(id uuid.UUID, post models.Post) error
	PublishPost(id uuid.UUID) error
	DeletePost(id uuid.UUID) error
}

type Usecase interface {
	CreatePost(post models.Post) error
	GetPostByID(id uuid.UUID) (models.Post, error)
	GetPosts() ([]models.Post, error)
	UpdatePost(id uuid.UUID, post models.Post) error
	PublishPost(id uuid.UUID) error
	DeletePost(id uuid.UUID) error
}
