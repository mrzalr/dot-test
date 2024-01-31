package server

import (
	"github.com/dot-test/internal/like/handler"
	likeRepository "github.com/dot-test/internal/like/repository/mysql"
	likeUsecase "github.com/dot-test/internal/like/usecase"
	postHandler "github.com/dot-test/internal/post/handler"
	postRepository "github.com/dot-test/internal/post/repository/mysql"
	postUsecase "github.com/dot-test/internal/post/usecase"
)

func (s *server) mapRoutes() {
	postRepo := postRepository.New(s.db)
	postUC := postUsecase.New(postRepo, s._redis)
	_postHandler := postHandler.New(postUC)

	likeRepo := likeRepository.New(s.db)
	likeUC := likeUsecase.New(likeRepo, postRepo)
	_likeHandler := handler.New(likeUC)

	v1 := s.app.Group("v1")
	postRoute := v1.Group("posts")

	_postHandler.MapRoutes(postRoute)
	_likeHandler.MapRoutes(postRoute)
}
