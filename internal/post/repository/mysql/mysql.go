package mysql

import (
	"github.com/dot-test/internal/models"
	"github.com/dot-test/internal/post"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) CreatePost(post models.Post) error {
	tx := r.db.Create(&post)
	return tx.Error
}

func New(db *gorm.DB) post.Repository {
	return &repository{
		db: db,
	}
}
