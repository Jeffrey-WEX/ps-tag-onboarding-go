package service

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/errormessage"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) GetUserById(userId string) (*model.User, *errormessage.ErrorMessage) {
	result := m.Called(userId)

	var r0 *model.User
	if result.Get(0) != nil {
		r0 = result.Get(0).(*model.User)
	}

	var r1 *errormessage.ErrorMessage
	if result.Get(1) != nil {
		r1 = result.Get(1).(*errormessage.ErrorMessage)
	}

	return r0, r1
}

func (m *UserServiceMock) CreateUser(newUser *model.User) (*model.User, *errormessage.ErrorMessage) {
	result := m.Called(newUser)

	var r0 *model.User
	if result.Get(0) != nil {
		r0 = result.Get(0).(*model.User)
	}

	var r1 *errormessage.ErrorMessage
	if result.Get(1) != nil {
		r1 = result.Get(1).(*errormessage.ErrorMessage)
	}

	return r0, r1
}
