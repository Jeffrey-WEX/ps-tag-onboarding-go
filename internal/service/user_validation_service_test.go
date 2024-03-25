package service

import (
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
	"github.com/stretchr/testify/suite"
)

type UserValidationServiceTestSuite struct {
	suite.Suite
	dbRepo                *repository.DbRepositoryMock
	userValidationService UserValidationService
}

func TestUserValidationService(t *testing.T) {
	suite.Run(t, new(UserValidationServiceTestSuite))
}

func (suite *UserValidationServiceTestSuite) SetupTest() {
	suite.dbRepo = &repository.DbRepositoryMock{}
	suite.userValidationService = NewUserValidationService(suite.dbRepo)
}

func (suite *UserValidationServiceTestSuite) TestValidateUserWithNoValidationErrors() {
	// Arrange
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	suite.dbRepo.On("FindUserByFirstLastName", user.FirstName, user.LastName).Return(model.User{})

	// Act
	valid, user := suite.userValidationService.ValidateUser(user)

	// Assert
	suite.Assert().True(valid)
	suite.Assert().Nil(user.ValidationErrors)
}

func (suite *UserValidationServiceTestSuite) TestVlidateUserWithAtLeastOneValidationError() {
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
	suite.Assert().False(valid)
	suite.Assert().NotNil(user.ValidationErrors)
	suite.Assert().Equal(3, len(user.ValidationErrors))
	suite.Assert().Contains(user.ValidationErrors, "User with the same first and last name already exists")
	suite.Assert().Contains(user.ValidationErrors, "User email must be properly formatted")
	suite.Assert().Contains(user.ValidationErrors, "User does not meet minimum age requirement")
}

func (suite *UserValidationServiceTestSuite) TestValidateAgeWithValidAge() {
	// Arrange
	user := model.User{
		Age: 25,
	}

	// Act
	validateAge(&user)

	// Arrange
	suite.Assert().Empty(user.ValidationErrors)
}

func (suite *UserValidationServiceTestSuite) TestValidateAgeWithAgeBelowMinimum() {
	// Arrange
	user := model.User{
		Age: 17,
	}

	// Act
	validateAge(&user)

	// Arrange
	suite.Assert().NotEmpty(user.ValidationErrors)
	suite.Assert().Equal(1, len(user.ValidationErrors))
	suite.Assert().Equal("User does not meet minimum age requirement", user.ValidationErrors[0])
}

func (suite *UserValidationServiceTestSuite) TestValidateEmailWithValidEmail() {
	// Arrange
	user := model.User{
		Email: "John.Doe@gmail.com",
	}

	// Act
	validateEmail(&user)

	// Arrange
	suite.Assert().Empty(user.ValidationErrors)
}

func (suite *UserValidationServiceTestSuite) TestValidateEmailWithMissingEmail() {
	// Arrange
	user := model.User{}

	// Act
	validateEmail(&user)

	// Arrange
	suite.Assert().NotEmpty(user.ValidationErrors)
	suite.Assert().Equal(1, len(user.ValidationErrors))
	suite.Assert().Equal("User email required", user.ValidationErrors[0])
}

func (suite *UserValidationServiceTestSuite) TestValidateEmailWithInvalidEmailFormat() {
	// Arrange
	user := model.User{
		Email: "JohnDoe.com",
	}

	// Act
	validateEmail(&user)

	// Arrange
	suite.Assert().NotEmpty(user.ValidationErrors)
	suite.Assert().Equal(1, len(user.ValidationErrors))
	suite.Assert().Equal("User email must be properly formatted", user.ValidationErrors[0])
}

func (suite *UserValidationServiceTestSuite) TestValidateNameWithValidName() {
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
	suite.Assert().Empty(user.ValidationErrors)
}

func (suite *UserValidationServiceTestSuite) TestValidateNameWithNameMissing() {
	// Arrange
	dbRepositoryMock := new(repository.DbRepositoryMock)
	userValidationService := NewUserValidationService(dbRepositoryMock)
	user := model.User{}

	// Act
	validateName(&user, userValidationService)

	// Arrange
	suite.Assert().NotEmpty(user.ValidationErrors)
	suite.Assert().Equal(1, len(user.ValidationErrors))
	suite.Assert().Equal("User first/last names required", user.ValidationErrors[0])
}

func (suite *UserValidationServiceTestSuite) TestValidateNameWithNameExistInDatabase() {
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
	suite.Assert().NotEmpty(user.ValidationErrors)
	suite.Assert().Equal(1, len(user.ValidationErrors))
	suite.Assert().Equal("User with the same first and last name already exists", user.ValidationErrors[0])
}
