package repository

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/stretchr/testify/mock"
)

type DbRepositoryMock struct {
	mock.Mock
}

func (m *DbRepositoryMock) GetUserById(userId string) (*model.User, error) {
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

func (m *DbRepositoryMock) CreateUser(newUser model.User) model.User {
	result := m.Called(newUser)

	return result.Get(0).(model.User)
}

func (m *DbRepositoryMock) FindUserByFirstLastName(firstName string, lastName string) model.User {
	result := m.Called(firstName, lastName)

	return result.Get(0).(model.User)
}
