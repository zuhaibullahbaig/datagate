package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(token string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		queryToken := c.Query("token")

		if strings.HasPrefix(c.Path(), "/static/") {
			return c.Next()
		}

		if queryToken != token {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized - Invalid or missing token",
			})
		}

		return c.Next()
	}
}
