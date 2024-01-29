package server

import (
	postHandler "github.com/dot-test/internal/post/handler"
	postRepository "github.com/dot-test/internal/post/repository/mysql"
	postUsecase "github.com/dot-test/internal/post/usecase"
)

func (s *server) mapRoutes() {
	postRepo := postRepository.New(s.db)
	postUC := postUsecase.New(postRepo)
	_postHandler := postHandler.New(postUC)

	v1 := s.app.Group("v1")
	postRoute := v1.Group("posts")

	_postHandler.MapRoutes(postRoute)
}
