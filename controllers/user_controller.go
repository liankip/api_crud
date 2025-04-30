package controllers

import (
	"api_crud/entities"
	"api_crud/usecases"
	"api_crud/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type UserController struct {
	UserUsecase *usecases.UserUsecase
}

func NewUserController(userUsecase *usecases.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

func (v *UserController) Signin(c *fiber.Ctx) error {
	var input entities.Signin

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

	user, err := v.UserUsecase.Signin(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
			Message: err.Error(),
			Data:    []string{},
		})
	}

	token, err := utils.GenerateJWT(uint(user.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Response{
			Message: "Failed to generate token",
			Data:    []string{},
		})
	}

	combinedData := fiber.Map{
		"user":  user,
		"token": token,
	}

	return c.JSON(entities.Response{
		Message: "Signin successful",
		Data:    combinedData,
	})
}

func (v *UserController) Signup(c *fiber.Ctx) error {
	var input entities.Signup

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

	if err := v.UserUsecase.Signup(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
			Message: err.Error(),
			Data:    []string{},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entities.Response{
		Message: "Signup successful",
		Data:    []string{},
	})
}
