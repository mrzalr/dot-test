package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	app *fiber.App
}

func New() *server {
	return &server{
		app: fiber.New(),
	}
}

func (s *server) Run(port string) error {
	return s.app.Listen(fmt.Sprintf(":%s", port))
}
