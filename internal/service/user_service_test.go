package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func setUpRepoAndService() (*UserService, *mocks.IUserRepository) {
	dbRepo := &mocks.IUserRepository{}
	userValidationService := NewUserValidationService()
	userService := NewService(dbRepo, *userValidationService)
	return userService, dbRepo
}

func TestGetUserById(t *testing.T) {
	t.Run("Get user sucessfully", func(t *testing.T) {
		// Arrange
		userService, dbRepo := setUpRepoAndService()
		user := &model.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@test.com",
			Age:       25,
		}

		dbRepo.On("GetUserById", "1").Return(user, nil)

		// Act
		result, _ := userService.GetUserById(user.ID)

		// Assert
		dbRepo.AssertCalled(t, "GetUserById", user.ID)
		assert.Equal(t, user, result)
	})

	t.Run("Get user return error when user not found", func(t *testing.T) {
		// Arrange
		userService, dbRepo := setUpRepoAndService()
		dbRepo.On("GetUserById", "1").Return(nil, errors.New(constant.ErrorUserNotFound))

		// Act
		result, err := userService.GetUserById("1")

		// Assert
		dbRepo.AssertCalled(t, "GetUserById", "1")
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, constant.ErrorUserNotFound, err.ErrorMessage)
	})

	t.Run("Get user return error when database returns error", func(t *testing.T) {
		// Arrange
		userService, dbRepo := setUpRepoAndService()
		dbRepo.On("GetUserById", "1").Return(nil, errors.New(constant.ErrorGettingUser))

		// Act
		result, err := userService.GetUserById("1")

		// Assert
		dbRepo.AssertCalled(t, "GetUserById", "1")
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, constant.ErrorGettingUser, err.ErrorMessage)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("Create user successfully", func(t *testing.T) {
		// Arrange
		userService, dbRepo := setUpRepoAndService()
		user := &model.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@test.com",
			Age:       25,
		}

		dbRepo.On("CreateUser", user).Return(user, nil)

		// Act
		newUser, errorMessage := userService.CreateUser(user)

		// Assert
		dbRepo.AssertCalled(t, "CreateUser", user)
		assert.Equal(t, user, newUser)
		assert.Empty(t, errorMessage)
	})

	t.Run("Create invalid user returns error", func(t *testing.T) {
		// Arrange
		userService, _ := setUpRepoAndService()
		user := &model.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoetest.com",
			Age:       16,
		}

		// Act
		newUser, errorMessage := userService.CreateUser(user)

		// Assert
		assert.Nil(t, newUser)
		assert.Equal(t, fmt.Sprintf("%s, %s", constant.ErrorAgeMinimum, constant.ErrorEmailInvalidFormat), errorMessage.ErrorMessage)
	})

	t.Run("Create existing user returns error", func(t *testing.T) {
		// Arrange
		userService, dbRepo := setUpRepoAndService()
		user := &model.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@test.com",
			Age:       24,
		}

		dbRepo.On("CreateUser", user).Return(nil, errors.New(constant.ErrorNameAlreadyExists))

		// Act
		newUser, errorMessage := userService.CreateUser(user)

		// Assert
		dbRepo.AssertCalled(t, "CreateUser", user)
		assert.Nil(t, newUser)
		assert.Equal(t, constant.ErrorNameAlreadyExists, errorMessage.ErrorMessage)
	})

	t.Run("Create user returns error when database returns error", func(t *testing.T) {
		// Arrange
		userService, dbRepo := setUpRepoAndService()
		user := &model.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@test.com",
			Age:       24,
		}

		dbRepo.On("CreateUser", user).Return(nil, errors.New(constant.ErrorCreatingUser))

		// Act
		newUser, errorMessage := userService.CreateUser(user)

		// Assert
		dbRepo.AssertCalled(t, "CreateUser", user)
		assert.Nil(t, newUser)
		assert.Equal(t, constant.ErrorCreatingUser, errorMessage.ErrorMessage)
	})
}
