package server

import (
	"fmt"

	errorhandler "github.com/dot-test/internal/middleware/error-handler"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type server struct {
	app *fiber.App
	db  *gorm.DB
}

func New(db *gorm.DB) *server {
	return &server{
		app: fiber.New(fiber.Config{
			ErrorHandler: errorhandler.CustomErrorHandler,
		}),
		db: db,
	}
}

func (s *server) Run(port string) error {
	s.mapRoutes()

	return s.app.Listen(fmt.Sprintf(":%s", port))
}
