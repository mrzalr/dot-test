package usecase

import (
	"github.com/dot-test/internal/models"
	"github.com/dot-test/internal/post"
)

type usecase struct {
	repository post.Repository
}

func (u *usecase) CreatePost(post models.Post) error {
	return u.repository.CreatePost(post)
}

func New(repository post.Repository) post.Usecase {
	return &usecase{
		repository: repository,
	}
}
