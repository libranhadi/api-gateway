package controller

import (
	"net/http"
	"service-user/helpers"
	"service-user/model"

	"service-user/service"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Auth(c *fiber.Ctx) error
}

type userControllerImpl struct {
	userService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &userControllerImpl{
		userService: userService,
	}
}

func (uc *userControllerImpl) Register(c *fiber.Ctx) error {
	requestBody := new(model.User)

	if err := c.BodyParser(requestBody); err != nil {
		return c.JSON(&helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid data request",
		})
	}

	err := uc.userService.Register(requestBody)

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
	} else {
		return c.JSON(&helpers.WebResponse{
			Code:   201,
			Status: "OK",
			Data:   requestBody.Email,
		})
	}
}

func (uc *userControllerImpl) Login(c *fiber.Ctx) error {
	requestBody := new(model.User)
	if err := c.BodyParser(requestBody); err != nil {
		return c.JSON(&helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid data request",
		})
	}
	var result model.User
	_, err := uc.userService.Login(requestBody)

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

	access_token := helpers.SignToken(requestBody.Email)

	return c.JSON(struct {
		Code        int
		Status      string
		AccessToken string
		Data        interface{}
	}{
		Code:        200,
		Status:      "OK",
		AccessToken: access_token,
		Data:        result,
	})
}

func (uc *userControllerImpl) Auth(c *fiber.Ctx) error {
	return c.JSON("OK")
}
