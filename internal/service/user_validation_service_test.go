package service

import (
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/stretchr/testify/assert"
)

func setUpUserValidationService() *UserValidationService {
	return NewUserValidationService()
}

func TestValidateUser(t *testing.T) {
	testCases := []struct {
		name               string
		user               model.User
		expectedErrors     []string
		expectedErrorCount int
	}{
		{
			name: "Validate user with no validation errors",
			user: model.User{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "JohnDoe@test.com",
				Age:       25,
			},
			expectedErrors:     nil,
			expectedErrorCount: 0,
		},
		{
			name: "Validate user with at least one validation error",
			user: model.User{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "JohnDoe.com",
				Age:       17,
			},
			expectedErrors: []string{
				"User email must be properly formatted",
				"User does not meet minimum age requirement",
			},
			expectedErrorCount: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			userValidationService := setUpUserValidationService()

			// Act
			validationErrors := userValidationService.ValidateUser(&tc.user)

			// Assert
			assert.Equal(t, tc.expectedErrorCount, len(validationErrors))
			for _, expectedError := range tc.expectedErrors {
				assert.Contains(t, validationErrors, expectedError)
			}
		})
	}
}

func TestValidateAge(t *testing.T) {
	testCases := []struct {
		name                 string
		age                  int
		expectedErrorMessage string
	}{
		{
			name:                 "Validate age with valid age",
			age:                  25,
			expectedErrorMessage: "",
		},
		{
			name:                 "Validate age with age below minimum",
			age:                  17,
			expectedErrorMessage: constant.ErrorAgeMinimum,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			user := model.User{
				Age: tc.age,
			}

			// Act
			validationError := validateAge(&user)

			// Assert
			assert.Equal(t, tc.expectedErrorMessage, validationError)
		})
	}
}

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		name                 string
		email                string
		expectedErrorMessage string
	}{
		{
			name:                 "Validate email with valid email",
			email:                "John.Doe@gmail.com",
			expectedErrorMessage: "",
		},
		{
			name:                 "Validate email with missing email",
			email:                "",
			expectedErrorMessage: constant.ErrorEmailRequired,
		},
		{
			name:                 "Validate email with invalid email format",
			email:                "JohnDoe.com",
			expectedErrorMessage: constant.ErrorEmailInvalidFormat,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			user := model.User{
				Email: tc.email,
			}

			// Act
			validationError := validateEmail(&user)

			// Assert
			assert.Equal(t, tc.expectedErrorMessage, validationError)
		})
	}
}
