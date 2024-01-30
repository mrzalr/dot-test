package handler

import "github.com/gofiber/fiber/v2"

func (h *handler) MapRoutes(r fiber.Router) {
	r.Post("/:postID/dislike", h.Dislike())
	r.Post("/:postID/like", h.Like())
}
