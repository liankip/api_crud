package main

import (
	"api_crud/controllers"
	"api_crud/middleware"
	"api_crud/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userRepository repository.UserRepository, userController *controllers.UserController, profileController *controllers.ProfileController) {

	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.Post("/signin", userController.Signin)
				auth.Post("/signup", userController.Signup)
			}

			profile := v1.Group("/profile", middleware.JWTAuthMiddleware)
			{
				profile.Get("/", middleware.RBACMiddleware(userRepository, "read"), profileController.CollectionProfile)
				profile.Get("/:id", middleware.RBACMiddleware(userRepository, "read"), profileController.DocumentProfile)

				management := profile.Group("/management/")
				{
					management.Post("/", middleware.RBACMiddleware(userRepository, "create"), profileController.CreateProfile)
					management.Put("/", middleware.RBACMiddleware(userRepository, "update"), profileController.UpdateProfile)
				}
			}
		}
	}
}
