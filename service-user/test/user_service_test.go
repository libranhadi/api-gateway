package service_user_test

import (
	"errors"
	"service-user/model"
	"service-user/repository"
	"service-user/service"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type UserServiceMockImpl struct {
	UserRepositoryMock *repository.UserRepositoryMock
}

func TestRegister_Success(t *testing.T) {
	mockRepo := new(repository.UserRepositoryMock)
	user := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}

	ctx := &fiber.Ctx{}

	mockRepo.On("FindUserByEmail", user.Email).Return(nil, errors.New("user not found"))
	mockRepo.On("Create", user).Return(nil)
	service := service.NewUserServiceImpl(mockRepo)

	err := service.Register(user, ctx)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
