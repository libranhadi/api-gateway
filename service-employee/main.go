package main

import (
	"fmt"
	"service-employee/config"
	"service-employee/controller"
	"service-employee/repository"
	"service-employee/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := config.GetPostgresDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi from service-employee")
	})

	empRepo := repository.NewEmployeeRepositoryImpl()
	empService := service.NewEmployeeService(empRepo, db)
	empController := controller.NewEmployeeControllerImpl(empService)
	app.Post("/employee", empController.CreateEmployee)

	port := 3002
	fmt.Printf("Service employee is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting Service employee: %v\n", err)
	}
}
