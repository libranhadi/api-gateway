package service

import (
	"net/http"
	"service-user/model"

	"service-user/repository"

	"service-user/helpers"
)

type UserService interface {
	Register(user *model.User) error
	Login(user *model.User) (*model.User, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: repository}
}

func (userService *UserServiceImpl) Register(user *model.User) error {

	if err := user.Validate(); err != nil {
		return &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
			Data:    user,
		}
	}

	userEmail, err := userService.userRepository.FindUserByEmail(user.Email)
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
	return userService.userRepository.Create(user)
}

func (userService *UserServiceImpl) Login(user *model.User) (*model.User, error) {

	if err := user.Validate(); err != nil {
		return user, &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
			Data:    user,
		}
	}

	userExist, err := userService.userRepository.FindUserByEmail(user.Email)
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
