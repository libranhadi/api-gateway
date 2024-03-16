package service

import (
	"database/sql"
	"net/http"
	"service-user/model"

	"service-user/repository"

	"service-user/helpers"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	Register(user *model.User, c *fiber.Ctx) error
	Login(user *model.User, c *fiber.Ctx) (*model.User, error)
}

type userServiceImpl struct {
	userRepository repository.UserRepository
	db             *sql.DB
}

func NewUserServiceImpl(repository repository.UserRepository, db *sql.DB) UserService {
	return &userServiceImpl{userRepository: repository, db: db}
}

func (userService *userServiceImpl) Register(user *model.User, c *fiber.Ctx) error {
	if err := c.BodyParser(user); err != nil {
		return c.JSON(&helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid data request",
		})
	}

	if err := user.Validate(); err != nil {
		return &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
			Data:    user,
		}
	}

	userEmail, err := userService.userRepository.FindUserByEmail(user.Email, userService.db)
	if err != nil {
		if err.Error() != "user not found" {
			return &helpers.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: err.Error(),
			}
		}
	}

	if userEmail != nil {
		return &helpers.WebResponse{
			Code:    http.StatusConflict,
			Status:  "Conflict",
			Message: "Email already exists",
		}
	}
	password := helpers.HashPassword([]byte(user.Password))
	user.Password = password
	return userService.userRepository.Create(user, userService.db)
}

func (userService *userServiceImpl) Login(user *model.User, c *fiber.Ctx) (*model.User, error) {
	if err := c.BodyParser(user); err != nil {
		return user, c.JSON(&helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid data request",
		})
	}

	if err := user.Validate(); err != nil {
		return user, &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
			Data:    user,
		}
	}

	userExist, err := userService.userRepository.FindUserByEmail(user.Email, userService.db)
	if err != nil {
		return userExist, &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
		}
	}

	if userExist == nil {
		return userExist, &helpers.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: err.Error(),
		}
	}

	checkPassword := helpers.ComparePassword([]byte(userExist.Password), []byte(user.Password))
	if !checkPassword {
		return userExist, &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Data:    userExist,
			Message: "Password doesn't match!",
		}
	}

	return userExist, nil
}
