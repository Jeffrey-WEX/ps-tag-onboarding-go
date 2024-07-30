package service

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/errormessage"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func setUpRepoAndService() (*UserService, *mocks.IDbRepository) {
	once = sync.Once{}
	dbRepo := new(mocks.IDbRepository)
	userValidationService := NewUserValidationService()
	userService := NewService(dbRepo, *userValidationService)
	return userService, dbRepo
}

func TestGetUserById(t *testing.T) {
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	userNotFoundErrorMessage := errormessage.ErrorMessage{
		ErrorMessage:    constant.ErrorUserNotFound,
		ErrorStatusCode: http.StatusNotFound,
	}

	errorGettingUserErrorMessage := errormessage.ErrorMessage{
		ErrorMessage:    constant.ErrorGettingUser,
		ErrorStatusCode: http.StatusInternalServerError,
	}

	testCases := []struct {
		name            string
		userID          string
		mockReturnUser  *model.User
		mockReturnError error
		expectedUser    *model.User
		expectedError   *errormessage.ErrorMessage
	}{
		{
			name:            "Get user successfully",
			userID:          "1",
			mockReturnUser:  &user,
			mockReturnError: nil,
			expectedUser:    &user,
			expectedError:   nil,
		},
		{
			name:            "Get user return error when user not found",
			userID:          "1",
			mockReturnUser:  nil,
			mockReturnError: fmt.Errorf("%s: %v", constant.ErrorUserNotFound, mongo.ErrNoDocuments),
			expectedUser:    nil,
			expectedError:   &userNotFoundErrorMessage,
		},
		{
			name:            "Get user return error when database returns error",
			userID:          "1",
			mockReturnUser:  nil,
			mockReturnError: errors.New(constant.ErrorGettingUser),
			expectedUser:    nil,
			expectedError:   &errorGettingUserErrorMessage,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			userService, dbRepo := setUpRepoAndService()
			dbRepo.On("GetUserById", tc.userID).Return(tc.mockReturnUser, tc.mockReturnError)

			// Act
			result, err := userService.GetUserById(tc.userID)

			// Assert
			dbRepo.AssertCalled(t, "GetUserById", tc.userID)
			assert.Equal(t, tc.expectedUser, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestCreateUser(t *testing.T) {
	validUser := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	invalidUser := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoetest.com",
		Age:       16,
	}

	// Test Create user successfully
	mockCreateUserFuncReturnUser := func(repo *mocks.IDbRepository) {
		repo.On("CreateUser", &validUser).Return(&validUser, nil)
	}

	// Test Create invalid user returns error
	invalidUserError := fmt.Sprintf("%s, %s", constant.ErrorAgeMinimum, constant.ErrorEmailInvalidFormat)
	invalidUserErrorMessage := errormessage.ErrorMessage{
		ErrorMessage:    invalidUserError,
		ErrorStatusCode: http.StatusBadRequest,
	}
	mockEmptyCreateUserFunc := func(repo *mocks.IDbRepository) {}

	// Test Create existing user returns error
	nameAlreadyExistsErrorMesage := errormessage.ErrorMessage{
		ErrorMessage:    constant.ErrorNameAlreadyExists,
		ErrorStatusCode: http.StatusBadRequest,
	}
	mockCreateUserFuncReturnNameAlreadyExistsError := func(repo *mocks.IDbRepository) {
		repo.On("CreateUser", &validUser).Return(nil, errors.New(constant.ErrorNameAlreadyExists))
	}

	// Test Create user returns error when database returns error
	internalServerErrorMessage := errormessage.ErrorMessage{
		ErrorMessage:    constant.ErrorCreatingUser,
		ErrorStatusCode: http.StatusInternalServerError,
	}
	mockCreateUserFuncReturnDatabaseError := func(repo *mocks.IDbRepository) {
		repo.On("CreateUser", &validUser).Return(nil, errors.New(constant.ErrorCreatingUser))
	}

	testCases := []struct {
		name                 string
		inputUser            *model.User
		mockCreateUserFunc   func(*mocks.IDbRepository)
		expectedUser         *model.User
		expectedErrorMessage *errormessage.ErrorMessage
	}{
		{
			name:                 "Create user successfully",
			inputUser:            &validUser,
			mockCreateUserFunc:   mockCreateUserFuncReturnUser,
			expectedUser:         &validUser,
			expectedErrorMessage: nil,
		},
		{
			name:                 "Create invalid user returns error",
			inputUser:            &invalidUser,
			mockCreateUserFunc:   mockEmptyCreateUserFunc,
			expectedUser:         nil,
			expectedErrorMessage: &invalidUserErrorMessage,
		},
		{
			name:                 "Create existing user returns error",
			inputUser:            &validUser,
			mockCreateUserFunc:   mockCreateUserFuncReturnNameAlreadyExistsError,
			expectedUser:         nil,
			expectedErrorMessage: &nameAlreadyExistsErrorMesage,
		},
		{
			name:                 "Create user returns error when database returns error",
			inputUser:            &validUser,
			mockCreateUserFunc:   mockCreateUserFuncReturnDatabaseError,
			expectedUser:         nil,
			expectedErrorMessage: &internalServerErrorMessage,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			userService, dbRepo := setUpRepoAndService()
			tc.mockCreateUserFunc(dbRepo)

			// Act
			newUser, errorMessage := userService.CreateUser(tc.inputUser)

			// Assert
			assert.Equal(t, tc.expectedUser, newUser)
			assert.Equal(t, tc.expectedErrorMessage, errorMessage)
		})
	}
}
