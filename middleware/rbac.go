package middleware

import (
	"api_crud/entities"
	"api_crud/repository"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func RBACMiddleware(userRepository repository.UserRepository, permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIDRaw := c.Locals("userID")
		if userIDRaw == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.Response{
				Message: "User not authenticated",
				Data:    nil,
			})
		}

		userID, ok := userIDRaw.(uint)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
				Message: "Invalid user ID format",
				Data:    nil,
			})
		}

		fmt.Println(userRepository.HasAccess(userID, permission))
		if !userRepository.HasAccess(userID, permission) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Permission denied",
			})
		}

		return c.Next()
	}
}
