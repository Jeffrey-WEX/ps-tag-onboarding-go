package service

import (
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestValidateUserWithNoValidationErrors(t *testing.T) {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)

	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	dbRepositoryMock.On("FindUserByFirstLastName", user.FirstName, user.LastName).Return(model.User{})

	// Act
	valid, user := userValidationService.ValidateUser(user)

	// Assert
	assert.True(t, valid)
	assert.Nil(t, user.ValidationErrors)
}

func TestVlidateUserWithAtLeastOneValidationError(t *testing.T) {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)

	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe.com",
		Age:       17,
	}

	dbRepositoryMock.On("FindUserByFirstLastName", user.FirstName, user.LastName).Return(user)

	// Act
	valid, user := userValidationService.ValidateUser(user)

	// Assert
	assert.False(t, valid)
	assert.NotNil(t, user.ValidationErrors)
	assert.Equal(t, 3, len(user.ValidationErrors))
	assert.Contains(t, user.ValidationErrors, "User with the same first and last name already exists")
	assert.Contains(t, user.ValidationErrors, "User email must be properly formatted")
	assert.Contains(t, user.ValidationErrors, "User does not meet minimum age requirement")
}

func TestValidateAgeWithValidAge(t *testing.T) {
	// Arrange
	user := model.User{
		Age: 25,
	}

	// Act
	validateAge(&user)

	// Arrange
	assert.Empty(t, user.ValidationErrors)
}

func TestValidateAgeWithAgeBelowMinimum(t *testing.T) {
	// Arrange
	user := model.User{
		Age: 17,
	}

	// Act
	validateAge(&user)

	// Arrange
	assert.NotEmpty(t, user.ValidationErrors)
	assert.Equal(t, 1, len(user.ValidationErrors))
	assert.Equal(t, "User does not meet minimum age requirement", user.ValidationErrors[0])
}

func TestValidateEmailWithValidEmail(t *testing.T) {
	// Arrange
	user := model.User{
		Email: "John.Doe@gmail.com",
	}

	// Act
	validateEmail(&user)

	// Arrange
	assert.Empty(t, user.ValidationErrors)
}

func TestValidateEmailWithMissingEmail(t *testing.T) {
	// Arrange
	user := model.User{}

	// Act
	validateEmail(&user)

	// Arrange
	assert.NotEmpty(t, user.ValidationErrors)
	assert.Equal(t, 1, len(user.ValidationErrors))
	assert.Equal(t, "User email required", user.ValidationErrors[0])
}

func TestValidateEmailWithInvalidEmailFormat(t *testing.T) {
	// Arrange
	user := model.User{
		Email: "JohnDoe.com",
	}

	// Act
	validateEmail(&user)

	// Arrange
	assert.NotEmpty(t, user.ValidationErrors)
	assert.Equal(t, 1, len(user.ValidationErrors))
	assert.Equal(t, "User email must be properly formatted", user.ValidationErrors[0])
}

func TestValidateNameWithValidName(t *testing.T) {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)
	user := model.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	dbRepositoryMock.On("FindUserByFirstLastName", user.FirstName, user.LastName).Return(model.User{})

	// Act
	validateName(&user, userValidationService)

	// Arrange
	assert.Empty(t, user.ValidationErrors)
}

func TestValidateNameWithNameMissing(t *testing.T) {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)
	user := model.User{}

	// Act
	validateName(&user, userValidationService)

	// Arrange
	assert.NotEmpty(t, user.ValidationErrors)
	assert.Equal(t, 1, len(user.ValidationErrors))
	assert.Equal(t, "User first/last names required", user.ValidationErrors[0])
}

func TestValidateNameWithNameExistInDatabase(t *testing.T) {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
	}

	dbRepositoryMock.On("FindUserByFirstLastName", user.FirstName, user.LastName).Return(user)

	// Act
	validateName(&user, userValidationService)

	// Arrange
	assert.NotEmpty(t, user.ValidationErrors)
	assert.Equal(t, 1, len(user.ValidationErrors))
	assert.Equal(t, "User with the same first and last name already exists", user.ValidationErrors[0])
}
