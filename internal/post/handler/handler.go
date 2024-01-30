package handler

import (
	"fmt"
	"net/http"

	"github.com/dot-test/internal/models"
	"github.com/dot-test/internal/post"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type handler struct {
	usecase post.Usecase
}

func (h *handler) CreatePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := models.NewPost()
		if err := c.BodyParser(&payload); err != nil {
			return post.ErrBadRequest
		}

		if err := h.usecase.CreatePost(payload); err != nil {
			return err
		}

		return c.JSON(models.Response{
			Status:  http.StatusCreated,
			Message: "created",
			Data:    map[string]any{"post_id": payload.ID},
			Error:   "",
		})
	}
}

func (h *handler) GetPostByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		_uuid, err := uuid.Parse(id)
		if err != nil {
			return post.ErrBadRequest
		}

		post, err := h.usecase.GetPostByID(_uuid)
		if err != nil {
			return err
		}

		return c.JSON(models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    post,
			Error:   "",
		})
	}
}

func (h *handler) GetPosts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := h.usecase.GetPosts()
		if err != nil {
			return err
		}

		return c.JSON(models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    posts,
			Error:   "",
		})
	}
}

func (h *handler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		_uuid, err := uuid.Parse(id)
		if err != nil {
			return post.ErrBadRequest
		}

		err = h.usecase.DeletePost(_uuid)
		if err != nil {
			return err
		}

		return c.JSON(models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    fmt.Sprintf("post with id %s already deleted", _uuid),
			Error:   "",
		})
	}
}

func (h *handler) PublishPost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		_uuid, err := uuid.Parse(id)
		if err != nil {
			return post.ErrBadRequest
		}

		err = h.usecase.PublishPost(_uuid)
		if err != nil {
			return err
		}

		return c.JSON(models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]any{"post_id": _uuid},
			Error:   "",
		})
	}
}

func (h *handler) UpdatePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		_uuid, err := uuid.Parse(id)
		if err != nil {
			return post.ErrBadRequest
		}

		payload := models.Post{}
		payload.ID = _uuid
		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		err = h.usecase.UpdatePost(_uuid, payload)
		if err != nil {
			return err
		}

		return c.JSON(models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]any{"post_id": _uuid},
			Error:   "",
		})
	}
}

func New(usecase post.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}
