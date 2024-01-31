package cacheflusher

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func CacheFlusher(_redis *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		if string(c.Request().Header.Method()) != http.MethodGet {
			_redis.FlushAll(c.Context())
			log.Println("Cache flushed")
		}

		return err
	}
}
