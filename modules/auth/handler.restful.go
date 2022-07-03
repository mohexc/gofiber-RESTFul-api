package auth

import "github.com/gofiber/fiber/v2"

func Route(router fiber.Group) {
	routerAdmin := router.Group("/admin/users")

	routerAdmin.Post("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"message": "Hello World!",
		})
	})

	routerAdmin.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"message": "Hello World!",
		})
	})

	routerAdmin.Get("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"message": "Hello World!",
		})
	})

	routerAdmin.Patch("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"message": "Hello World!",
		})
	})

	routerAdmin.Delete("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"message": "Hello World!",
		})
	})

}
