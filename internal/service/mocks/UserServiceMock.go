package mocks

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) GetAllUsers() []model.User {
	result := m.Called()

	return result.Get(0).([]model.User)
}

func (m *UserServiceMock) GetUserById(userId string) (*model.User, error) {
	result := m.Called(userId)

	var r0 *model.User
	if result.Get(0) != nil {
		r0 = result.Get(0).(*model.User)
	}

	var r1 error
	if result.Get(1) != nil {
		r1 = result.Get(1).(error)
	}

	return r0, r1
}

func (m *UserServiceMock) CreateUser(newUser model.User) model.User {
	result := m.Called(newUser)

	return result.Get(0).(model.User)
}
