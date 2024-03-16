package repository

import (
	"errors"
	"service-user/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (repoMock *UserRepositoryMock) FindUserByEmail(email string) (*model.User, error) {
	args := repoMock.Mock.Called(email)
	if args.Get(0) == nil {
		return nil, errors.New("user not found")
	}
	user := args.Get(0).(model.User)
	return &user, nil
}

func (repoMock *UserRepositoryMock) Create(user *model.User) error {
	args := repoMock.Mock.Called(user)
	return args.Error(0)
}
