package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	fmt.Println("AdminMiddleware")
	return c.Next()
}
