package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	router         *gin.Engine
	userService    *service.UserServiceMock
	userController UserController
}

func TestUserController(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.router = gin.Default()
	suite.userService = new(service.UserServiceMock)
	suite.userController = NewController(suite.userService)

}

func (suite *UserControllerTestSuite) TestControllerGetUserById() {
	// Arrange
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	suite.userService.On("GetUserById", "1").Return(&user, nil)

	// Act
	suite.router.GET("/users/:id", suite.userController.GetUserById)
	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	var bodyResponse model.User
	json.Unmarshal(w.Body.Bytes(), &bodyResponse)

	// Assert
	suite.userService.AssertCalled(suite.T(), "GetUserById", "1")
	suite.Assert().Equal(http.StatusOK, w.Code)
	suite.Assert().Equal(user, bodyResponse)
}

func (suite *UserControllerTestSuite) TestUserControllerCreaterUser() {
	// Arrange
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	suite.userService.On("CreateUser", user).Return(user)

	// Act
	suite.router.POST("/users", suite.userController.CreateUser)
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	var bodyResponse model.User
	json.Unmarshal(w.Body.Bytes(), &bodyResponse)

	// Assert
	suite.userService.AssertCalled(suite.T(), "CreateUser", user)
	suite.Assert().Equal(http.StatusCreated, w.Code)
	suite.Assert().Equal(user, bodyResponse)
}
