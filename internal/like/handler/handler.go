package handler

import (
	"net/http"

	"github.com/dot-test/internal/like"
	"github.com/dot-test/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type handler struct {
	usecase like.Usecase
}

func (h *handler) Dislike() fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID := c.Params("postID")
		postUUID, err := uuid.Parse(postID)
		if err != nil {
			return like.ErrBadRequest
		}

		payload := models.Like{}
		payload.PostID = postUUID
		if err := c.BodyParser(&payload); err != nil {
			return like.ErrBadRequest
		}

		if err := h.usecase.Dislike(payload); err != nil {
			return err
		}

		return c.JSON(models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data: map[string]any{
				"user_id": payload.UserID,
				"post_id": postUUID,
			},
			Error: "",
		})
	}
}

func (h *handler) Like() fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID := c.Params("postID")
		postUUID, err := uuid.Parse(postID)
		if err != nil {
			return like.ErrBadRequest
		}

		payload := models.Like{}
		payload.PostID = postUUID
		if err := c.BodyParser(&payload); err != nil {
			return like.ErrBadRequest
		}

		if err := h.usecase.Like(payload); err != nil {
			return err
		}

		return c.JSON(models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data: map[string]any{
				"user_id": payload.UserID,
				"post_id": postUUID,
			},
			Error: "",
		})
	}
}

func New(usecase like.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}
