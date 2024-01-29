package handler

import "github.com/gofiber/fiber/v2"

func (h *handler) MapRoutes(r fiber.Router) {
	r.Post("", h.CreatePost())
}
