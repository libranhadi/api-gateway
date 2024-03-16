package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"service-user/helpers"
	"service-user/repository"

	"github.com/gofiber/fiber/v2"
)

type Auth interface {
	Authentication(c *fiber.Ctx) error
}

type authImpl struct {
	userRepository repository.UserRepository
	db             *sql.DB
}

func NewAuthImpl(repository repository.UserRepository, db *sql.DB) Auth {
	return &authImpl{userRepository: repository, db: db}
}

func (auth *authImpl) Authentication(c *fiber.Ctx) error {
	access_token := c.Get("access_token")

	if len(access_token) == 0 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}

	checkToken, err := helpers.VerifyToken(access_token)

	if err != nil {
		return c.Status(401).SendString("Invalid token: Failed to verify token")
	}

	fmt.Println(checkToken, "CEKKKK", checkToken["email"])

	email, ok := checkToken["email"].(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(&helpers.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Sorry, something went wrong please ",
		})
	}

	user, err := auth.userRepository.FindUserByEmail(email, auth.db)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&helpers.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
	}

	if user == nil {
		return c.Status(http.StatusNotFound).JSON(&helpers.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: err.Error(),
		})
	}

	// Set user data in context for future use
	c.Locals("user", user)

	// Continue processing if user is found
	return c.Next()
}
