package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	repomocks "github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	dbRepo                *repomocks.IUserRepository
	userValidationService UserValidationService
	userService           UserService
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.dbRepo = &repomocks.IUserRepository{}
	suite.userValidationService = NewUserValidationService()
	suite.userService = NewService(suite.dbRepo, suite.userValidationService)

}

func (suite *UserServiceTestSuite) TestGetUserUsingExistingId_ReturnsUser() {
	// Arrange
	user := &model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	suite.dbRepo.On("GetUserById", "1").Return(user, nil)

	// Act
	result, _ := suite.userService.GetUserById(user.ID)

	// Assert
	suite.dbRepo.AssertCalled(suite.T(), "GetUserById", user.ID)
	suite.Assert().Equal(user, result)
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
	user := &model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	suite.dbRepo.On("CreateUser", user).Return(user, nil)

	// Act
	newUser, errorMessage := suite.userService.CreateUser(user)

	// Assert
	suite.dbRepo.AssertCalled(suite.T(), "CreateUser", user)
	suite.Assert().Equal(user, newUser)
	suite.Assert().Empty(errorMessage)
}

func (suite *UserServiceTestSuite) TestCreateInvalidUser_ReturnsError() {
	// Arrange
	user := &model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoetest.com",
		Age:       16,
	}

	// Act
	newUser, errorMessage := suite.userService.CreateUser(user)

	// Assert
	suite.Assert().Nil(newUser)
	suite.Assert().Equal(fmt.Sprintf("%s, %s", constant.ErrorAgeMinimum, constant.ErrorEmailInvalidFormat), errorMessage.ErrorMessage)
}

func (suite *UserServiceTestSuite) TestCreateExistingUser_ReturnsError() {
	// Arrange
	user := &model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       24,
	}

	suite.dbRepo.On("CreateUser", user).Return(nil, errors.New(constant.ErrorNameAlreadyExists))

	// Act
	newUser, errorMessage := suite.userService.CreateUser(user)

	// Assert
	suite.dbRepo.AssertCalled(suite.T(), "CreateUser", user)
	suite.Assert().Nil(newUser)
	suite.Assert().Equal(constant.ErrorNameAlreadyExists, errorMessage.ErrorMessage)
}
