package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dot-test/internal/models"
	"github.com/dot-test/internal/post"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type usecase struct {
	repository post.Repository
	_redis     *redis.Client
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

func (u *usecase) GetPopularPost() ([]models.Post, error) {
	popularPost, err := u._redis.Get(context.Background(), "popular_post").Result()
	if err != nil {
		if err == redis.Nil {
			posts, err := u.repository.GetPopularPost()
			if err != nil {
				return []models.Post{}, err
			}

			jsonPosts, err := json.Marshal(posts)
			if err != nil {
				return []models.Post{}, err
			}

			_, err = u._redis.Set(context.Background(), "popular_post", jsonPosts, 0).Result()
			if err != nil {
				return []models.Post{}, err
			}

			log.Println("Get popular posts from db")
			return posts, nil
		} else {
			return []models.Post{}, err
		}
	}

	posts := []models.Post{}
	err = json.Unmarshal([]byte(popularPost), &posts)
	if err != nil {
		return []models.Post{}, err
	}

	log.Println("Get popular posts from redis")
	return posts, nil
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

func New(repository post.Repository, _redis *redis.Client) post.Usecase {
	return &usecase{
		repository: repository,
		_redis:     _redis,
	}
}
