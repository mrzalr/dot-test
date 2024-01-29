package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type server struct {
	app *fiber.App
	db  *gorm.DB
}

func New(db *gorm.DB) *server {
	return &server{
		app: fiber.New(),
		db:  db,
	}
}

func (s *server) Run(port string) error {
	return s.app.Listen(fmt.Sprintf(":%s", port))
}
