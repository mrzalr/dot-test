package like

import (
	"github.com/dot-test/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(trx *gorm.DB) (*gorm.DB, Repository)
	Like(like models.Like) error
	Dislike(like models.Like) error
}
type Usecase interface {
	Like(like models.Like) error
	Dislike(like models.Like) error
}
