package service

import (
	"errors"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetUserUsingExistingId_ReturnsUser(t *testing.T) {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)
	userService := NewService(dbRepositoryMock, userValidationService)

	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	dbRepositoryMock.On("GetUserById", "1").Return(&user, nil)

	// Act
	result, _ := userService.GetUserById(user.ID)

	// Assert
	dbRepositoryMock.AssertCalled(t, "GetUserById", "1")
	assert.Equal(t, user, *result)
}

func TestGetUserUsingNonExistingId_ReturnsError(t *testing.T) {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)
	userService := NewService(dbRepositoryMock, userValidationService)

	dbRepositoryMock.On("GetUserById", "1").Return(nil, errors.New("user not found"))

	// Act
	result, err := userService.GetUserById("1")

	// Assert
	dbRepositoryMock.AssertCalled(t, "GetUserById", "1")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestGetAllUsers_ReturnsUsers(t *testing.T) {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)
	userService := NewService(dbRepositoryMock, userValidationService)

	user1 := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	user2 := model.User{
		ID:        "2",
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "JaneDoe@test.com",
		Age:       23,
	}

	users := []model.User{user1, user2}

	dbRepositoryMock.On("GetAllUsers").Return(users)

	// Act
	result := userService.GetAllUsers()

	// Assert
	dbRepositoryMock.AssertCalled(t, "GetAllUsers")
	assert.Equal(t, users, result)
}

func TestCreateValidUser_ReturnsUser(t *testing.T) {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)
	userService := NewService(dbRepositoryMock, userValidationService)

	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	dbRepositoryMock.On("CreateUser", user).Return(user)
	dbRepositoryMock.On("FindUserByFirstLastName", user.FirstName, user.LastName).Return(model.User{})

	// Act
	result := userService.CreateUser(user)

	// Assert
	dbRepositoryMock.AssertCalled(t, "FindUserByFirstLastName", user.FirstName, user.LastName)
	dbRepositoryMock.AssertCalled(t, "CreateUser", user)
	assert.Equal(t, user, result)
	assert.Empty(t, user.ValidationErrors)
}

func TestCreateInvalidUser_ReturnsError(t *testing.T) {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)
	userService := NewService(dbRepositoryMock, userValidationService)

	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       24,
	}

	dbRepositoryMock.On("FindUserByFirstLastName", user.FirstName, user.LastName).Return(user)

	// Act
	result := userService.CreateUser(user)

	// Assert
	dbRepositoryMock.AssertNotCalled(t, "CreateUser", user)
	dbRepositoryMock.AssertCalled(t, "FindUserByFirstLastName", user.FirstName, user.LastName)
	assert.NotEmpty(t, result.ValidationErrors)
	assert.Equal(t, "User with the same first and last name already exists", result.ValidationErrors[0])
}
