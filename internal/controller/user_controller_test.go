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
	"github.com/stretchr/testify/assert"
)

func TestUserControllerGetAllUsers(t *testing.T) {
	// Arrange
	router := gin.Default()
	userServiceMock := new(service.UserServiceMock)
	userController := NewController(userServiceMock)
	users := []model.User{{ID: "1"}, {ID: "2"}}
	userServiceMock.On("GetAllUsers").Return(users, nil)

	// Act
	router.GET("/users", userController.GetAllUsers)
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var bodyResponse []model.User
	json.Unmarshal(w.Body.Bytes(), &bodyResponse)

	// Assert
	userServiceMock.AssertCalled(t, "GetAllUsers")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, users, bodyResponse)
}

func TestControllerGetUserById(t *testing.T) {
	// Arrange
	router := gin.Default()
	userServiceMock := new(service.UserServiceMock)
	userController := NewController(userServiceMock)

	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	userServiceMock.On("GetUserById", "1").Return(&user, nil)

	// Act
	router.GET("/users/:id", userController.GetUserById)
	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var bodyResponse model.User
	json.Unmarshal(w.Body.Bytes(), &bodyResponse)

	// Assert
	userServiceMock.AssertCalled(t, "GetUserById", "1")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user, bodyResponse)
}

func TestUserControllerCreaterUser(t *testing.T) {
	// Arrange
	router := gin.Default()
	userServiceMock := new(service.UserServiceMock)
	userController := NewController(userServiceMock)

	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	userServiceMock.On("CreateUser", user).Return(user)

	// Act
	router.POST("/users", userController.CreateUser)
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var bodyResponse model.User
	json.Unmarshal(w.Body.Bytes(), &bodyResponse)

	// Assert
	userServiceMock.AssertCalled(t, "CreateUser", user)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, user, bodyResponse)
}
