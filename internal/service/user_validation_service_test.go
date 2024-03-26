package service

import (
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/stretchr/testify/suite"
)

type UserValidationServiceTestSuite struct {
	suite.Suite
	userValidationService UserValidationService
}

func TestUserValidationService(t *testing.T) {
	suite.Run(t, new(UserValidationServiceTestSuite))
}

func (suite *UserValidationServiceTestSuite) SetupTest() {
	suite.userValidationService = NewUserValidationService()
}

func (suite *UserValidationServiceTestSuite) TestValidateUserWithNoValidationErrors() {
	// Arrange
	user := model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	// Act
	validationErrors := suite.userValidationService.ValidateUser(&user)

	// Assert
	suite.Assert().Nil(validationErrors)
}

func (suite *UserValidationServiceTestSuite) TestVlidateUserWithAtLeastOneValidationError() {
	// Arrange

	user := model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe.com",
		Age:       17,
	}

	// Act
	validationErrors := suite.userValidationService.ValidateUser(&user)

	// Assert
	suite.Assert().NotNil(validationErrors)
	suite.Assert().Equal(2, len(validationErrors))
	suite.Assert().Contains(validationErrors, "User email must be properly formatted")
	suite.Assert().Contains(validationErrors, "User does not meet minimum age requirement")
}

func (suite *UserValidationServiceTestSuite) TestValidateAgeWithValidAge() {
	// Arrange
	user := model.User{
		Age: 25,
	}

	// Act
	validationError := validateAge(&user)

	// Arrange
	suite.Assert().Empty(validationError)
}

func (suite *UserValidationServiceTestSuite) TestValidateAgeWithAgeBelowMinimum() {
	// Arrange
	user := model.User{
		Age: 17,
	}

	// Act
	validationError := validateAge(&user)

	// Arrange
	suite.Assert().NotEmpty(validationError)
	suite.Assert().Equal(constant.ErrorAgeMinimum, validationError)
}

func (suite *UserValidationServiceTestSuite) TestValidateEmailWithValidEmail() {
	// Arrange
	user := model.User{
		Email: "John.Doe@gmail.com",
	}

	// Act
	validationError := validateEmail(&user)

	// Arrange
	suite.Assert().Empty(validationError)
}

func (suite *UserValidationServiceTestSuite) TestValidateEmailWithMissingEmail() {
	// Arrange
	user := model.User{}

	// Act
	validationError := validateEmail(&user)

	// Arrange
	suite.Assert().NotEmpty(validationError)
	suite.Assert().Equal(constant.ErrorEmailRequired, validationError)
}

func (suite *UserValidationServiceTestSuite) TestValidateEmailWithInvalidEmailFormat() {
	// Arrange
	user := model.User{
		Email: "JohnDoe.com",
	}

	// Act
	validationError := validateEmail(&user)

	// Arrange
	suite.Assert().NotEmpty(validationError)
	suite.Assert().Equal(constant.ErrorEmailInvalidFormat, validationError)
}
