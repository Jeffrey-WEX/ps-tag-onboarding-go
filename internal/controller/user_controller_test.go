package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/errormessage"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func getGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func TestControllerGetUserById(t *testing.T) {

	t.Run("Get user sucessfully", func(t *testing.T) {
		// Arrange
		userServiceMock := new(mocks.IService)
		userController := NewController(userServiceMock)
		w := httptest.NewRecorder()
		ctx := getGinContext(w)

		user := model.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@test.com",
			Age:       25,
		}

		userServiceMock.On("GetUserById", "1").Return(&user, nil)
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		// Act
		userController.GetUserById(ctx)

		// Assert
		assert.True(t, userServiceMock.AssertCalled(t, "GetUserById", "1"))
		assert.Equal(t, http.StatusOK, w.Code)

		expectedBody, err := json.MarshalIndent(user, "", "    ")
		if err != nil {
			t.Fatalf("Error marshaling JSON: %v", err)
		}
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("Get user failed with error message returned", func(t *testing.T) {
		// Arrange
		userServiceMock := new(mocks.IService)
		userController := NewController(userServiceMock)
		w := httptest.NewRecorder()
		ctx := getGinContext(w)

		errorMessage := errormessage.ErrorMessage{
			ErrorStatusCode: http.StatusBadRequest,
			ErrorMessage:    constant.ErrorAgeMinimum,
		}

		userServiceMock.On("GetUserById", "1").Return(nil, &errorMessage)
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		// Act
		userController.GetUserById(ctx)

		// Assert
		assert.True(t, userServiceMock.AssertCalled(t, "GetUserById", "1"))
		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedBody, err := json.MarshalIndent(errorMessage, "", "    ")
		if err != nil {
			t.Fatalf("Error marshaling JSON: %v", err)
		}
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

}

func TestUserControllerCreateUser(t *testing.T) {

	t.Run("Create user successfully", func(t *testing.T) {
		// Arrange
		userServiceMock := new(mocks.IService)
		userController := NewController(userServiceMock)
		w := httptest.NewRecorder()
		ctx := getGinContext(w)

		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@test.com",
			Age:       25,
		}

		userServiceMock.On("CreateUser", &user).Return(&user, nil)
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		body, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("Error marshaling JSON: %v", err)
		}

		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// Act
		userController.CreateUser(ctx)

		// Assert
		assert.True(t, userServiceMock.AssertCalled(t, "CreateUser", &user))
		assert.Equal(t, http.StatusCreated, w.Code)

		expectedBody, err := json.MarshalIndent(user, "", "    ")
		if err != nil {
			t.Fatalf("Error marshaling JSON: %v", err)
		}
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("Create user failed with bad JSON", func(t *testing.T) {
		// Arrange
		userServiceMock := new(mocks.IService)
		userController := NewController(userServiceMock)
		w := httptest.NewRecorder()
		ctx := getGinContext(w)

		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(bytes.NewBufferString(""))

		// Act
		userController.CreateUser(ctx)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		expectedBody, err := json.MarshalIndent(gin.H{"status_code": http.StatusBadRequest, "message": constant.ErrorInvalidUserObject}, "", "    ")
		if err != nil {
			t.Fatalf("Error marshaling JSON: %v", err)
		}
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("Create user failed with error message returned", func(t *testing.T) {
		// Arrange
		userServiceMock := new(mocks.IService)
		userController := NewController(userServiceMock)
		w := httptest.NewRecorder()
		ctx := getGinContext(w)

		errorMessage := errormessage.ErrorMessage{
			ErrorStatusCode: http.StatusBadRequest,
			ErrorMessage:    constant.ErrorEmailInvalidFormat,
		}

		userServiceMock.On("CreateUser", &model.User{}).Return(nil, &errorMessage)
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		body, err := json.Marshal(model.User{})
		if err != nil {
			t.Fatalf("Error marshaling JSON: %v", err)
		}

		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// Act
		userController.CreateUser(ctx)

		// Assert
		assert.True(t, userServiceMock.AssertCalled(t, "CreateUser", &model.User{}))
		assert.Equal(t, http.StatusBadRequest, w.Code)

		expectedBody, err := json.MarshalIndent(errorMessage, "", "    ")
		if err != nil {
			t.Fatalf("Error marshaling JSON: %v", err)
		}
		assert.Equal(t, string(expectedBody), w.Body.String())

	})
}
