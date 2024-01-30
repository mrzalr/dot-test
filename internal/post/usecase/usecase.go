package usecase

import (
	"fmt"

	"github.com/dot-test/internal/models"
	"github.com/dot-test/internal/post"
	"github.com/google/uuid"
)

type usecase struct {
	repository post.Repository
}

func (u *usecase) DeletePost(id uuid.UUID) error {
	_, err := u.repository.GetPostByID(id)
	if err != nil {
		return err
	}

	return u.repository.DeletePost(id)
}

func (u *usecase) GetPostByID(id uuid.UUID) (models.Post, error) {
	return u.repository.GetPostByID(id)
}

func (u *usecase) GetPosts() ([]models.Post, error) {
	return u.repository.GetPosts()
}

func (u *usecase) PublishPost(id uuid.UUID) error {
	found, err := u.repository.GetPostByID(id)
	if err != nil {
		return err
	}

	if found.IsPublished {
		return fmt.Errorf("the post is already published")
	}

	return u.repository.PublishPost(id)
}

func (u *usecase) UpdatePost(id uuid.UUID, post models.Post) error {
	_, err := u.repository.GetPostByID(id)
	if err != nil {
		return err
	}

	return u.repository.UpdatePost(id, post)
}

func (u *usecase) CreatePost(post models.Post) error {
	return u.repository.CreatePost(post)
}

func New(repository post.Repository) post.Usecase {
	return &usecase{
		repository: repository,
	}
}
