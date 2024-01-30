package mysql

import (
	"github.com/dot-test/internal/like"
	"github.com/dot-test/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) WithTrx(trx *gorm.DB) (*gorm.DB, like.Repository) {
	if trx == nil {
		trx = r.db.Begin()
	}

	return trx, &repository{
		db: trx,
	}
}

func (r *repository) Dislike(like models.Like) error {
	tx := r.db.Where("user_id = ? AND post_id = ?", like.UserID, like.PostID).Delete(&models.Like{})
	return tx.Error
}

func (r *repository) Like(like models.Like) error {
	tx := r.db.Create(like)
	return tx.Error
}

func New(db *gorm.DB) like.Repository {
	return &repository{
		db: db,
	}
}
