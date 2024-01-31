package mysql

import (
	"github.com/dot-test/internal/models"
	"github.com/dot-test/internal/post"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) WithTrx(trx *gorm.DB) (*gorm.DB, post.Repository) {
	if trx == nil {
		trx = r.db.Begin()
	}

	return trx, &repository{
		db: trx,
	}
}

func (r *repository) GetPopularPost() ([]models.Post, error) {
	popularPost := []models.Post{}
	tx := r.db.Where("is_published = ?", true).
		Order("like_count DESC").Limit(10).Find(&popularPost)
	return popularPost, tx.Error
}

func (r *repository) UpdateLikesCount(id uuid.UUID, likeCount int) error {
	tx := r.db.Model(&models.Post{}).Where("id = ?", id).Update("like_count", likeCount)
	return tx.Error
}

func (r *repository) DeletePost(id uuid.UUID) error {
	tx := r.db.Where("id = ?", id).Delete(&models.Post{})
	return tx.Error
}

func (r *repository) PublishPost(id uuid.UUID) error {
	tx := r.db.Model(&models.Post{}).Where("id = ?", id).Update("is_published", true)
	return tx.Error
}

func (r *repository) UpdatePost(id uuid.UUID, post models.Post) error {
	tx := r.db.Model(&models.Post{}).Where("id = ?", id).Updates(post)
	return tx.Error
}

func (r *repository) CreatePost(post models.Post) error {
	tx := r.db.Create(&post)
	return tx.Error
}

func (r *repository) GetPostByID(id uuid.UUID) (models.Post, error) {
	post := models.Post{}
	tx := r.db.Where("id = ?", id).First(&post)
	return post, tx.Error
}

func (r *repository) GetPosts() ([]models.Post, error) {
	posts := []models.Post{}
	tx := r.db.Where("is_published = ?", true).Find(&posts)
	return posts, tx.Error
}

func New(db *gorm.DB) post.Repository {
	return &repository{
		db: db,
	}
}
