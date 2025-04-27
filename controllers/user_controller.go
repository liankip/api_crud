package controllers

import (
	"api_crud/usecases"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserUsecase *usecases.UserUsecase
}

func NewUserController(userUsecase *usecases.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

func (v *UserController) Signup(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
