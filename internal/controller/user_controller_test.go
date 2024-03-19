package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserControllerGetAllUsers(t *testing.T) {
	router := gin.Default()
	userServiceMock := new(mocks.UserServiceMock)
	userController := NewController(userServiceMock)

	// Set up the mock to return a slice of users
	users := []model.User{{ID: "1"}, {ID: "2"}}
	userServiceMock.On("GetAllUsers").Return(users, nil)

	// Call the method
	router.GET("/users", userController.GetAllUsers)
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the mock method was called
	userServiceMock.AssertCalled(t, "GetAllUsers")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "1", users[0].ID)
	assert.Equal(t, "2", users[1].ID)
}

func TestControllerGetUserById(t *testing.T) {
	router := gin.Default()
	userServiceMock := new(mocks.UserServiceMock)
	userController := NewController(userServiceMock)

	// Set up the mock to return a user
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	userServiceMock.On("GetUserById", "1").Return(&user, nil)

	// Call the method
	router.GET("/users/:id", userController.GetUserById)
	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the mock method was called
	userServiceMock.AssertCalled(t, "GetUserById", "1")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "1", user.ID)
}

func TestUserControllerCreaterUser(t *testing.T) {
	router := gin.Default()
	userServiceMock := new(mocks.UserServiceMock)
	userController := NewController(userServiceMock)

	// Set up the mock to return a user
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}
	userServiceMock.On("CreateUser", user).Return(user)

	// Call the method
	router.POST("/users", userController.CreateUser)
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the mock method was called
	userServiceMock.AssertCalled(t, "CreateUser", user)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "1", user.ID)
}
