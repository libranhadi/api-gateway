package controller

import (
	"fmt"
	"net/http"
	"service-employee/helpers"
	"service-employee/model"
	"service-employee/service"

	"github.com/gofiber/fiber/v2"
)

var user_uri string = "http://service-user:3001/user"

type EmployeeController interface {
	CreateEmployee(c *fiber.Ctx) error
}

type employeeControllerImpl struct {
	employeeService service.EmployeeService
}

func NewEmployeeControllerImpl(empService service.EmployeeService) EmployeeController {
	return &employeeControllerImpl{
		employeeService: empService,
	}
}

func (empController *employeeControllerImpl) CreateEmployee(c *fiber.Ctx) error {
	requestBody := new(model.Employee)

	if err := c.BodyParser(requestBody); err != nil {
		return c.JSON(&helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid data request",
		})
	}

	access_token := c.Get("access_token")

	resp, err := empController.employeeService.ConnectUserService(user_uri, access_token)

	if err != nil {
		webResponse, ok := err.(*helpers.WebResponse)
		if ok {
			return c.Status(webResponse.Code).JSON(webResponse)
		} else {
			return c.Status(http.StatusInternalServerError).JSON(&helpers.WebResponse{
				Code:    http.StatusInternalServerError,
				Status:  "Internal server error",
				Message: "Sorry, something went wrong please try again later!",
			})
		}
	}
	defer resp.Body.Close()

	// Print the response
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)

	if resp.Status != "200 OK" {
		return c.Status(401).JSON(helpers.WebResponse{
			Code:    401,
			Status:  "401 unautorhized",
			Message: "Please login for create employee",
		})
	}

	errCreate := empController.employeeService.CreateEmployee(requestBody)
	if errCreate != nil {
		webResponse, ok := errCreate.(*helpers.WebResponse)
		if ok {
			return c.Status(webResponse.Code).JSON(webResponse)
		} else {
			return c.Status(http.StatusInternalServerError).JSON(&helpers.WebResponse{
				Code:    http.StatusInternalServerError,
				Status:  "Internal server error",
				Message: "Sorry, something went wrong please try again later!",
			})
		}
	} else {
		return c.JSON(helpers.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   requestBody,
		})
	}
}
