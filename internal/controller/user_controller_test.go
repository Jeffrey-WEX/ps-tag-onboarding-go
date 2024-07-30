package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/errormessage"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func getGinContext(w *httptest.ResponseRecorder) *gin.Context {
	once = sync.Once{}
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func TestController_GetUserById(t *testing.T) {
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	errorMessage := errormessage.ErrorMessage{
		ErrorStatusCode: http.StatusBadRequest,
		ErrorMessage:    constant.ErrorAgeMinimum,
	}

	testCases := []struct {
		name                   string
		userID                 string
		mockReturnUser         *model.User
		mockReturnError        *errormessage.ErrorMessage
		expectedStatusCodeCode int
		expectedResponseBody   interface{}
	}{
		{
			name:                   "Get user successfully",
			userID:                 "1",
			mockReturnUser:         &user,
			mockReturnError:        nil,
			expectedStatusCodeCode: http.StatusOK,
			expectedResponseBody:   user,
		},
		{
			name:                   "Get user failed with error message returned",
			userID:                 "1",
			mockReturnUser:         nil,
			mockReturnError:        &errorMessage,
			expectedStatusCodeCode: http.StatusBadRequest,
			expectedResponseBody:   errorMessage,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			userServiceMock := new(mocks.IUserService)
			userController := NewController(userServiceMock)
			w := httptest.NewRecorder()
			ctx := getGinContext(w)

			userServiceMock.On("GetUserById", tc.userID).Return(tc.mockReturnUser, tc.mockReturnError)
			ctx.Request.Method = "GET"
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: tc.userID})

			// Act
			userController.GetUserById(ctx)

			// Assert
			assert.True(t, userServiceMock.AssertCalled(t, "GetUserById", tc.userID))
			assert.Equal(t, tc.expectedStatusCodeCode, w.Code)

			expectedBody, _ := json.MarshalIndent(tc.expectedResponseBody, "", "    ")
			assert.Equal(t, string(expectedBody), w.Body.String())
		})
	}
}

func TestUserController_CreateUser(t *testing.T) {
	user := model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	// Test create user succesfully
	mockCreateUserFuncReturnUser := func(mock *mocks.IUserService) {
		mock.On("CreateUser", &user).Return(&user, nil)
	}

	// Test create user failed with bad JSON
	badJSONBody := gin.H{
		"status_code": http.StatusBadRequest,
		"message":     constant.ErrorInvalidUserObject,
	}
	mockEmptyCreateUserFunc := func(mock *mocks.IUserService) {}

	// Test create user failed with error message returned
	badRequestErrorMessage := errormessage.ErrorMessage{
		ErrorStatusCode: http.StatusBadRequest,
		ErrorMessage:    constant.ErrorEmailInvalidFormat,
	}
	mockCreateUserFuncReturnError := func(mock *mocks.IUserService) {
		mock.On("CreateUser", &model.User{}).Return(nil, &badRequestErrorMessage)
	}

	testCases := []struct {
		name               string
		mockCreateUserFunc func(*mocks.IUserService)
		inputBody          interface{}
		expectedStatusCode int
		expectedBody       interface{}
	}{
		{
			name:               "Create user successfully",
			mockCreateUserFunc: mockCreateUserFuncReturnUser,
			inputBody:          user,
			expectedStatusCode: http.StatusCreated,
			expectedBody:       user,
		},
		{
			name:               "Create user failed with bad JSON",
			mockCreateUserFunc: mockEmptyCreateUserFunc,
			inputBody:          "",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       badJSONBody,
		},
		{
			name:               "Create user failed with error message returned",
			mockCreateUserFunc: mockCreateUserFuncReturnError,
			inputBody:          model.User{},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       badRequestErrorMessage,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			userServiceMock := new(mocks.IUserService)
			tc.mockCreateUserFunc(userServiceMock)
			userController := NewController(userServiceMock)
			w := httptest.NewRecorder()
			ctx := getGinContext(w)

			ctx.Request.Method = "POST"
			ctx.Request.Header.Set("Content-Type", "application/json")
			body, _ := json.Marshal(tc.inputBody)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

			// Act
			userController.CreateUser(ctx)

			// Assert
			assert.Equal(t, tc.expectedStatusCode, w.Code)

			expectedBody, _ := json.MarshalIndent(tc.expectedBody, "", "    ")
			assert.Equal(t, string(expectedBody), w.Body.String())
		})
	}
}
