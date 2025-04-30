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
