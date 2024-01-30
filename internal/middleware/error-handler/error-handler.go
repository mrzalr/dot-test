package errorhandler

import (
	"net/http"

	"github.com/dot-test/internal/like"
	"github.com/dot-test/internal/models"
	"github.com/dot-test/internal/post"
	"github.com/gofiber/fiber/v2"
)

var errorMap = map[error]int{
	post.ErrBadRequest: http.StatusBadRequest,
	like.ErrBadRequest: http.StatusBadRequest,
}

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	code, ok := errorMap[err]
	if !ok {
		code = http.StatusInternalServerError
	}

	return c.JSON(models.Response{
		Status:  code,
		Message: "error",
		Data:    nil,
		Error:   err.Error(),
	})
}
