package service

import (
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/stretchr/testify/assert"
)

func setUpUserValidationService() UserValidationService {
	return NewUserValidationService()
}

func TestValidateUser(t *testing.T) {

	t.Run("Validate user with no validation errors", func(t *testing.T) {
		// Arrange
		userValidationService := setUpUserValidationService()
		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@test.com",
			Age:       25,
		}

		// Act
		validationErrors := userValidationService.ValidateUser(&user)

		// Assert
		assert.Nil(t, validationErrors)
	})

	t.Run("Validate user with at least one validation error", func(t *testing.T) {
		// Arrange
		userValidationService := setUpUserValidationService()
		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe.com",
			Age:       17,
		}

		// Act
		validationErrors := userValidationService.ValidateUser(&user)

		// Assert
		assert.NotNil(t, validationErrors)
		assert.Equal(t, 2, len(validationErrors))
		assert.Contains(t, validationErrors, "User email must be properly formatted")
		assert.Contains(t, validationErrors, "User does not meet minimum age requirement")
	})
}

func TestValidateAge(t *testing.T) {
	t.Run("Validate age with valid age", func(t *testing.T) {
		// Arrange
		user := model.User{
			Age: 25,
		}

		// Act
		validationError := validateAge(&user)

		// Assert
		assert.Empty(t, validationError)
	})

	t.Run("Validate age with age below minimum", func(t *testing.T) {
		// Arrange
		user := model.User{
			Age: 17,
		}

		// Act
		validationError := validateAge(&user)

		// Assert
		assert.NotEmpty(t, validationError)
		assert.Equal(t, constant.ErrorAgeMinimum, validationError)
	})

}

func TestValidateEmail(t *testing.T) {
	t.Run("Validate email with valid email", func(t *testing.T) {
		// Arrange
		user := model.User{
			Email: "John.Doe@gmail.com",
		}

		// Act
		validationError := validateEmail(&user)

		// Assert
		assert.Empty(t, validationError)
	})

	t.Run("Validate email with missing email", func(t *testing.T) {
		// Arrange
		user := model.User{}

		// Act
		validationError := validateEmail(&user)

		// Assert
		assert.NotEmpty(t, validationError)
		assert.Equal(t, constant.ErrorEmailRequired, validationError)
	})

	t.Run("Validate email with invalid email format", func(t *testing.T) {
		// Arrange
		user := model.User{
			Email: "JohnDoe.com",
		}

		// Act
		validationError := validateEmail(&user)

		// Assert
		assert.NotEmpty(t, validationError)
		assert.Equal(t, constant.ErrorEmailInvalidFormat, validationError)
	})
}
