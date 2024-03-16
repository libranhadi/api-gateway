package main

import (
	"fmt"

	"service-user/config"
	"service-user/controller"
	"service-user/middleware"
	"service-user/repository"
	"service-user/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.GetPostgresDB()
	userRepo := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(userRepo, db)
	userController := controller.NewUserControllerImpl(userService)
	middleware := middleware.NewAuthImpl(userRepo, db)
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi from service-user")
	})
	app.Post("/user/register", userController.Register)
	app.Post("/user/login", userController.Login)
	app.Get("/user/auth", middleware.Authentication, userController.Auth)

	port := 3001
	fmt.Printf("Service user is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting Service user: %v\n", err)
	}
}
