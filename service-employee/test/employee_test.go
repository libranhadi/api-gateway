package test

import (
	"errors"
	"service-employee/model"
	"service-employee/repository"
	"service-employee/service"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestCreateServiceWithMockRepository_Success(t *testing.T) {
	mockRepo := new(repository.EmployeeRepositoryMock)
	employee := &model.Employee{
		Name: "test",
	}

	mockRepo.On("Create", employee).Return(nil)
	serviceMock := service.NewEmployeeServiceImpl(mockRepo)
	err := serviceMock.CreateEmployee(employee)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateServiceWithMockRepository_FailedValidation(t *testing.T) {
	mockRepo := new(repository.EmployeeRepositoryMock)
	employee := &model.Employee{
		Name: "Test",
	}
	objError := errors.New("name field is required")
	mockRepo.On("Create", employee).Return(objError)
	serviceMock := service.NewEmployeeServiceImpl(mockRepo)
	err := serviceMock.CreateEmployee(employee)
	assert.EqualError(t, err, objError.Error())
	mockRepo.AssertExpectations(t)
}
