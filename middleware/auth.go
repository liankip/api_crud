package middleware

import (
	"api_crud/entities"
	"api_crud/utils"
	"github.com/gofiber/fiber/v2"
)

func JWTAuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(entities.Response{
			Message: "Missing token",
			Data:    []string{},
		})
	}

	userID, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(entities.Response{
			Message: "Invalid token",
			Data:    []string{},
		})
	}

	c.Locals("userID", userID)
	return c.Next()
}
