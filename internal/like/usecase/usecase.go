package usecase

import (
	"github.com/dot-test/internal/like"
	"github.com/dot-test/internal/models"
	"github.com/dot-test/internal/post"
)

type usecase struct {
	repository     like.Repository
	postRepository post.Repository
}

func (u *usecase) Dislike(like models.Like) error {
	trx, trxRepo := u.repository.WithTrx(nil)
	err := trxRepo.Dislike(like)
	if err != nil {
		trx.Rollback()
		return err
	}

	_, trxPostRepo := u.postRepository.WithTrx(trx)
	post, err := trxPostRepo.GetPostByID(like.PostID)
	if err != nil {
		trx.Rollback()
		return err
	}

	updatedLikeCount := post.LikeCount - 1
	err = trxPostRepo.UpdateLikesCount(like.PostID, updatedLikeCount)
	if err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit().Error
}

func (u *usecase) Like(like models.Like) error {
	trx, trxRepo := u.repository.WithTrx(nil)
	err := trxRepo.Like(like)
	if err != nil {
		trx.Rollback()
		return err
	}

	_, trxPostRepo := u.postRepository.WithTrx(trx)
	post, err := trxPostRepo.GetPostByID(like.PostID)
	if err != nil {
		trx.Rollback()
		return err
	}

	updatedLikeCount := post.LikeCount + 1
	err = trxPostRepo.UpdateLikesCount(like.PostID, updatedLikeCount)
	if err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit().Error
}

func New(repository like.Repository, postRepository post.Repository) like.Usecase {
	return &usecase{
		repository:     repository,
		postRepository: postRepository,
	}
}
