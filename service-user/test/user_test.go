package test

import (
	"errors"
	"service-user/helpers"
	"service-user/model"
	"service-user/repository"
	"service-user/service"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestRegisterWithMockRepository_Success(t *testing.T) {
	mockRepo := new(repository.UserRepositoryMock)
	user := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}

	mockRepo.On("FindUserByEmail", user.Email).Return(nil, errors.New("user not found"))
	mockRepo.On("Create", user).Return(nil)
	serviceMock := service.NewUserServiceImpl(mockRepo)
	err := serviceMock.Register(user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegisterWithMockRepository_Failed(t *testing.T) {
	mockRepo := new(repository.UserRepositoryMock)
	user := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}
	objError := errors.New("email already exist")
	mockRepo.On("FindUserByEmail", user.Email).Return(nil, objError)
	mockRepo.On("Create", user).Return(objError)
	serviceMock := service.NewUserServiceImpl(mockRepo)
	err := serviceMock.Register(user)
	assert.EqualError(t, err, objError.Error())
	mockRepo.AssertExpectations(t)
}

func TestLoginWithMockRepository_Success(t *testing.T) {
	mockRepo := new(repository.UserRepositoryMock)

	userExist := &model.User{
		Email:    "email_test_1@example.com",
		Password: helpers.HashPassword([]byte("password")),
	}

	mockRepo.On("FindUserByEmail", userExist.Email).Return(userExist, nil)

	userLogin := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}

	serviceMock := service.NewUserServiceImpl(mockRepo)
	_, err := serviceMock.Login(userLogin)
	assert.Nil(t, err)
	assert.Equal(t, userExist.Email, userLogin.Email)
	mockRepo.AssertExpectations(t)
}

func TestLoginWithMockRepository_Failed(t *testing.T) {
	mockRepo := new(repository.UserRepositoryMock)

	userLogin := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}
	objError := errors.New("user not found")
	mockRepo.On("FindUserByEmail", userLogin.Email).Return(nil, objError)

	serviceMock := service.NewUserServiceImpl(mockRepo)
	_, err := serviceMock.Login(userLogin)
	assert.EqualError(t, err, objError.Error())
	mockRepo.AssertExpectations(t)

}

func TestLoginWithMockRepository_Failed_PasswordNotMatch(t *testing.T) {
	mockRepo := new(repository.UserRepositoryMock)

	userExist := &model.User{
		Email:    "email_test_1@example.com",
		Password: helpers.HashPassword([]byte("test")),
	}

	objError := errors.New("Password doesn't match!")

	mockRepo.On("FindUserByEmail", userExist.Email).Return(userExist, objError)

	userLogin := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}

	serviceMock := service.NewUserServiceImpl(mockRepo)
	_, err := serviceMock.Login(userLogin)
	assert.EqualError(t, err, objError.Error())
	mockRepo.AssertExpectations(t)
}
