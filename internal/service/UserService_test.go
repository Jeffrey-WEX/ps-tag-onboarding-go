package service

import (
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	userServiceMock := new(mocks.UserServiceMock)
	userServiceMock.On("GetAllUsers").Return([]model.User{{ID: "1"}, {ID: "2"}}, nil)

	users := userServiceMock.GetAllUsers()

	assert.Len(t, users, 2)
	assert.Equal(t, "1", users[0].ID)
	assert.Equal(t, "2", users[1].ID)

	userServiceMock.AssertExpectations(t)
}

func TestGetUserById(t *testing.T) {
	userServiceMock := new(mocks.UserServiceMock)
	userServiceMock.On("GetUserById", "1")

	user, err := userServiceMock.GetUserById("1")

	assert.NoError(t, err)
	assert.Equal(t, "1", user.ID)

	userServiceMock.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	userServiceMock := new(mocks.UserServiceMock)
	userServiceMock.On("CreateUser", model.User{ID: "1"}).Return(model.User{ID: "1"}, nil)

	user := userServiceMock.CreateUser(model.User{ID: "1"})

	assert.Equal(t, "1", user.ID)

	userServiceMock.AssertExpectations(t)
}
