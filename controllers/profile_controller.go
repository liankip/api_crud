package controllers

import (
	"api_crud/entities"
	"api_crud/usecases"
	"github.com/gofiber/fiber/v2"
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
