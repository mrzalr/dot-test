package handler

import "github.com/gofiber/fiber/v2"

func (h *handler) MapRoutes(r fiber.Router) {
	r.Post("", h.CreatePost())
	r.Get("/popular", h.GetPopularPost())
	r.Get("/:id", h.GetPostByID())
	r.Get("", h.GetPosts())
	r.Put("/:id", h.UpdatePost())
	r.Patch("/:id/publish", h.PublishPost())
	r.Delete("/:id", h.Delete())
}
