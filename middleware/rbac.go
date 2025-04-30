package middleware

import (
	"api_crud/repository"
	"github.com/gofiber/fiber/v2"
)

func RBACMiddleware(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Locals("id")
		if id == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User ID not found",
			})
		}

		userID, ok := id.(int)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID type",
			})
		}

		if !repository.UserRepository.HasAccess(userID, permission) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Permission denied",
			})
		}

		return c.Next()
	}
}
