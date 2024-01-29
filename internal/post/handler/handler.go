package handler

import (
	"github.com/dot-test/internal/models"
	"github.com/dot-test/internal/post"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase post.Usecase
}

func (h *handler) CreatePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := models.NewPost()
		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		if err := h.usecase.CreatePost(payload); err != nil {
			return err
		}

		return c.JSON(models.NewResponseCreated(nil))
	}
}

func New(usecase post.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}
