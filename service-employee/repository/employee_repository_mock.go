package repository

import (
	"service-employee/model"

	"github.com/stretchr/testify/mock"
)

type EmployeeRepositoryMock struct {
	mock.Mock
}

func (repoMock *EmployeeRepositoryMock) Create(employee *model.Employee) error {
	args := repoMock.Mock.Called(employee)
	return args.Error(0)
}
