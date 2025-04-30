package controllers

import (
	"api_crud/entities"
	"api_crud/usecases"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ProfileController struct {
	ProfileUsecase *usecases.ProfileUsecase
}

func NewProfileController(profileUsecase *usecases.ProfileUsecase) *ProfileController {
	return &ProfileController{ProfileUsecase: profileUsecase}
}

func (v *ProfileController) CollectionProfile(c *fiber.Ctx) error {
	profile, _ := v.ProfileUsecase.CollectionProfile()

	return c.JSON(entities.Response{
		Message: "Collection Profile successful",
		Data:    profile,
	})
}

func (v *ProfileController) DocumentProfile(c *fiber.Ctx) error {
	idParam := c.Params("id")

	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	id := uint(idUint64)

	profile, _ := v.ProfileUsecase.DocumentProfile(id)

	return c.JSON(entities.Response{
		Message: "Document Profile successful",
		Data:    profile,
	})
}

func (v *ProfileController) CreateProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	var input entities.CreateProfile

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
			Message: "Invalid input format",
			Data:    err.Error(),
		})
	}

	if err := validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
			Message: "Validation failed",
			Data:    err.Error(),
		})
	}

	profile, err := v.ProfileUsecase.CreateProfile(entities.CreateProfile{
		UserID:    userID.(uint),
		Bio:       input.Bio,
		AvatarUrl: input.AvatarUrl,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
			Message: "Profile Create failed",
			Data:    err.Error(),
		})
	}

	return c.JSON(entities.Response{
		Message: "Profile Create successful",
		Data:    profile,
	})
}
