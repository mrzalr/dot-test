package server

import (
	"fmt"

	cacheflusher "github.com/dot-test/internal/middleware/cache-flusher"
	errorhandler "github.com/dot-test/internal/middleware/error-handler"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type server struct {
	app    *fiber.App
	db     *gorm.DB
	_redis *redis.Client
}

func New(db *gorm.DB, _redis *redis.Client) *server {
	return &server{
		app:    fiber.New(fiber.Config{ErrorHandler: errorhandler.CustomErrorHandler}),
		db:     db,
		_redis: _redis,
	}
}

func (s *server) Run(port string) error {
	s.app.Use(cacheflusher.CacheFlusher(s._redis))
	s.mapRoutes()

	return s.app.Listen(fmt.Sprintf(":%s", port))
}
