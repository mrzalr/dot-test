package post

import (
	"github.com/dot-test/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(trx *gorm.DB) (*gorm.DB, Repository)
	CreatePost(post models.Post) error
	GetPostByID(id uuid.UUID) (models.Post, error)
	GetPosts() ([]models.Post, error)
	GetPopularPost() ([]models.Post, error)
	UpdatePost(id uuid.UUID, post models.Post) error
	PublishPost(id uuid.UUID) error
	DeletePost(id uuid.UUID) error
	UpdateLikesCount(id uuid.UUID, likeCount int) error
}

type Usecase interface {
	CreatePost(post models.Post) error
	GetPostByID(id uuid.UUID) (models.Post, error)
	GetPosts() ([]models.Post, error)
	GetPopularPost() ([]models.Post, error)
	UpdatePost(id uuid.UUID, post models.Post) error
	PublishPost(id uuid.UUID) error
	DeletePost(id uuid.UUID) error
}
