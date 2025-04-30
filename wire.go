// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"api_crud/controllers"
	"api_crud/db"
	"api_crud/repository"
	"api_crud/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	connection, err := db.ConnectDB("root:password@tcp(127.0.0.1:3306)/db_api_crud?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	return connection, nil
}

var RepositorySet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewProfileRepository,
)

var UsecaseSet = wire.NewSet(
	usecases.NewUserUsecase,
	usecases.NewProfileUsecase,
)

var ControllerSet = wire.NewSet(
	controllers.NewUserController,
	controllers.NewProfileController,
)

func InitializeApplication() (*fiber.App, error) {
	wire.Build(
		InitializeDB,
		RepositorySet,
		UsecaseSet,
		ControllerSet,
		NewApp,
	)
	return &fiber.App{}, nil
}

func NewApp(userRepository repository.UserRepository, userController *controllers.UserController, profileController *controllers.ProfileController) *fiber.App {
	app := fiber.New()
	SetupRoutes(app, userRepository, userController, profileController)
	return app
}
