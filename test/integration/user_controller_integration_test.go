package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/errormessage"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository/database"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/route"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func cleanUpDb(db *mongo.Database) {
	databaseName := os.Getenv("DATABASE_NAME")
	collection := db.Collection(databaseName)
	_, err := collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
}

func setUpAppAndDb(user *model.User, addUser bool) (*gin.Engine, *mongo.Database, repository.IDbRepository) {
	err := godotenv.Load("../../variables.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var db = database.NewDatabase()
	var userRepository = repository.NewRepository(db)
	var userValidator = service.NewUserValidationService()
	var userService = service.NewService(userRepository, *userValidator)
	var userController = controller.NewController(userService)
	var routes = route.NewRoutes(userController)
	router := gin.Default()
	routes.InitializeRouter(router)

	if addUser {
		userRepository.CreateUser(user)
	}

	return router, db, userRepository
}

func TestUserControllerIntegration_GetUserById(t *testing.T) {
	userOne := model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohDoe@test.com",
		Age:       25,
	}

	emptyErrorMessage := errormessage.ErrorMessage{}

	userNotFoundErrorMessage := errormessage.ErrorMessage{
		ErrorStatusCode: http.StatusNotFound,
		ErrorMessage:    constant.ErrorUserNotFound,
	}

	nonExistingUser := model.User{
		ID: "1",
	}

	testCases := []struct {
		name                 string
		user                 *model.User
		addUser              bool
		expectedStatusCode   int
		expectedErrorMessage errormessage.ErrorMessage
	}{
		{
			name:                 "Return user when finding an existing user",
			user:                 &userOne,
			addUser:              true,
			expectedStatusCode:   http.StatusOK,
			expectedErrorMessage: emptyErrorMessage,
		},
		{
			name:                 "Return not found when finding a non-existing user",
			user:                 &nonExistingUser,
			addUser:              false,
			expectedStatusCode:   http.StatusNotFound,
			expectedErrorMessage: userNotFoundErrorMessage,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			router, db, _ := setUpAppAndDb(tc.user, tc.addUser)
			url := fmt.Sprintf("/v1/users/%s", tc.user.ID)
			req, _ := http.NewRequest("GET", url, nil)
			w := httptest.NewRecorder()

			// Act
			router.ServeHTTP(w, req)

			// Assert
			var errorMessage errormessage.ErrorMessage
			_ = json.Unmarshal(w.Body.Bytes(), &errorMessage)

			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.Equal(t, tc.expectedErrorMessage, errorMessage)

			t.Cleanup(func() {
				cleanUpDb(db)
			})
		})
	}
}

func TestUserControllerIntegration_CreateUser(t *testing.T) {
	validUser := model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@test.com",
		Age:       25,
	}

	invalidUser := model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe",
		Age:       13,
	}

	emptyErrorMessage := errormessage.ErrorMessage{}

	invalidUserErrorMessage := errormessage.ErrorMessage{
		ErrorStatusCode: http.StatusBadRequest,
		ErrorMessage:    "User does not meet minimum age requirement, User email must be properly formatted",
	}

	userAlreadyExistsErrorMessage := errormessage.ErrorMessage{
		ErrorStatusCode: http.StatusBadRequest,
		ErrorMessage:    constant.ErrorNameAlreadyExists,
	}

	testCases := []struct {
		name                 string
		user                 *model.User
		addUser              bool
		expectedStatusCode   int
		expectedErrorMessage errormessage.ErrorMessage
	}{
		{
			name:                 "Creating a valid user",
			user:                 &validUser,
			addUser:              false,
			expectedStatusCode:   http.StatusCreated,
			expectedErrorMessage: emptyErrorMessage,
		},
		{
			name:                 "Creating an invalid user",
			user:                 &invalidUser,
			addUser:              false,
			expectedStatusCode:   http.StatusBadRequest,
			expectedErrorMessage: invalidUserErrorMessage,
		},
		{
			name:                 "Creating a user with an existing name",
			user:                 &validUser,
			addUser:              true,
			expectedStatusCode:   http.StatusBadRequest,
			expectedErrorMessage: userAlreadyExistsErrorMessage,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			router, db, _ := setUpAppAndDb(tc.user, tc.addUser)
			jsonValue, _ := json.Marshal(tc.user)
			req, _ := http.NewRequest("POST", "/v1/users", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()

			// Act
			router.ServeHTTP(w, req)

			// Assert
			var errorMessage errormessage.ErrorMessage
			_ = json.Unmarshal(w.Body.Bytes(), &errorMessage)

			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.Equal(t, tc.expectedErrorMessage, errorMessage)

			t.Cleanup(func() {
				cleanUpDb(db)
			})
		})
	}
}
