package service

import (
	"errors"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	dbRepo                *repository.DbRepositoryMock
	userValidationService UserValidationService
	userService           UserService
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.dbRepo = &repository.DbRepositoryMock{}
	suite.userValidationService = NewUserValidationService(suite.dbRepo)
	suite.userService = NewService(suite.dbRepo, suite.userValidationService)

}

func (suite *UserServiceTestSuite) TestGetUserUsingExistingId_ReturnsUser() {
	// Arrange
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	suite.dbRepo.On("GetUserById", "1").Return(&user, nil)

	// Act
	result, _ := suite.userService.GetUserById(user.ID)

	// Assert
	suite.dbRepo.AssertCalled(suite.T(), "GetUserById", user.ID)
	suite.Assert().Equal(user, *result)
}

func (suite *UserServiceTestSuite) TestGetUserUsingNonExistingId_ReturnsError() {
	// Arrange
	suite.dbRepo.On("GetUserById", "1").Return(nil, errors.New("user not found"))

	// Act
	result, err := suite.userService.GetUserById("1")

	// Assert
	suite.dbRepo.AssertCalled(suite.T(), "GetUserById", "1")
	suite.Assert().Nil(result)
	suite.Assert().NotNil(err)
}

func (suite *UserServiceTestSuite) TestCreateValidUser_ReturnsUser() {
	// Arrange
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	suite.dbRepo.On("CreateUser", user).Return(user)
	suite.dbRepo.On("FindUserByFirstLastName", user.FirstName, user.LastName).Return(model.User{})

	// Act
	result := suite.userService.CreateUser(user)

	// Assert
	suite.dbRepo.AssertCalled(suite.T(), "FindUserByFirstLastName", user.FirstName, user.LastName)
	suite.dbRepo.AssertCalled(suite.T(), "CreateUser", user)
	suite.Assert().Equal(user, result)
	suite.Assert().Empty(user.ValidationErrors)
}

func (suite *UserServiceTestSuite) TestCreateInvalidUser_ReturnsError() {
	// Arrange
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       24,
	}

	suite.dbRepo.On("FindUserByFirstLastName", user.FirstName, user.LastName).Return(user)

	// Act
	result := suite.userService.CreateUser(user)

	// Assert
	suite.dbRepo.AssertNotCalled(suite.T(), "CreateUser", user)
	suite.dbRepo.AssertCalled(suite.T(), "FindUserByFirstLastName", user.FirstName, user.LastName)
	suite.Assert().NotEmpty(result.ValidationErrors)
	suite.Assert().Equal("User with the same first and last name already exists", result.ValidationErrors[0])
}
