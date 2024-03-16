package service

import (
	"database/sql"
	"net/http"
	"service-employee/model"

	"service-employee/helpers"
	"service-employee/repository"

	"github.com/gofiber/fiber/v2"
)

type EmployeeService interface {
	Create(employee *model.Employee, c *fiber.Ctx) error
	ConnectUserService(user_uri string, c *fiber.Ctx) (*http.Response, error)
}

type employeeServiceImpl struct {
	employeeRepository repository.EmployeeRepository
	db                 *sql.DB
}

func NewEmployeeService(repository repository.EmployeeRepository, db *sql.DB) EmployeeService {
	return &employeeServiceImpl{employeeRepository: repository, db: db}
}

func (employeeService *employeeServiceImpl) Create(employee *model.Employee, c *fiber.Ctx) error {

	if err := c.BodyParser(employee); err != nil {
		return c.JSON(&helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid data request",
		})
	}

	if err := employee.Validate(); err != nil {
		return &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
			Data:    employee,
		}
	}

	return employeeService.employeeRepository.Create(employee, employeeService.db)
}

func (employeeService *employeeServiceImpl) ConnectUserService(user_uri string, c *fiber.Ctx) (*http.Response, error) {

	access_token := c.Get("access_token")
	if len(access_token) == 0 {
		return nil, &helpers.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "401 unauthorized",
			Message: "Invalid token: Access token missing",
		}
	}

	req, err := http.NewRequest("GET", user_uri+"/auth", nil)
	if err != nil {
		return nil, &helpers.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Internal server error",
			Message: "Failed to create request",
		}
		// fmt.Println("Error creating request:", err)
		// panic(err)
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", access_token)

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, &helpers.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Internal server error",
			Message: "Failed to perform request",
		}
		// fmt.Println("Error sending request:", err)
		// panic(err)
	}
	return resp, nil
}
