package auth

import "github.com/gofiber/fiber/v2"

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("authToken")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Error{
				Message: "Token tidak ditemukan, silakan login terlebih dahulu!",
			})
		}

		c.Locals("token", token)

		return c.Next()
	}
}
