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

func setUpAppAndDb() (*gin.Engine, *mongo.Database, repository.IUserRepository) {
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

	return router, db, userRepository
}

func TestUserControllerIntegration(t *testing.T) {

	t.Run("Return not found when finding a non-existing user", func(t *testing.T) {
		router, db, _ := setUpAppAndDb()

		req, _ := http.NewRequest("GET", "/v1/users/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var errorMessage errormessage.ErrorMessage
		err := json.Unmarshal(w.Body.Bytes(), &errorMessage)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, http.StatusNotFound, errorMessage.ErrorStatusCode)
		assert.Equal(t, constant.ErrorUserNotFound, errorMessage.ErrorMessage)

		t.Cleanup(func() {
			cleanUpDb(db)
		})
	})

	t.Run("Return user when finding an existing user", func(t *testing.T) {
		router, db, userRepository := setUpAppAndDb()
		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohDoe@test.com",
			Age:       25,
		}
		userRepository.CreateUser(&user)
		url := fmt.Sprintf("/v1/users/%s", user.ID)
		req, _ := http.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var bodyResponse model.User
		err := json.Unmarshal(w.Body.Bytes(), &bodyResponse)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, user.FirstName, bodyResponse.FirstName)
		assert.Equal(t, user.LastName, bodyResponse.LastName)
		assert.Equal(t, user.Email, bodyResponse.Email)
		assert.Equal(t, user.Age, bodyResponse.Age)

		t.Cleanup(func() {
			cleanUpDb(db)
		})

	})

	t.Run("Creating a valid user", func(t *testing.T) {
		router, db, _ := setUpAppAndDb()
		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@test.com",
			Age:       25,
		}

		jsonValue, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("Error marshaling JSON: %v", err)
		}
		req, _ := http.NewRequest("POST", "/v1/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var bodyResponse model.User
		err = json.Unmarshal(w.Body.Bytes(), &bodyResponse)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, user.FirstName, bodyResponse.FirstName)

		t.Cleanup(func() {
			cleanUpDb(db)
		})

	})

	t.Run("Creating an invalid user", func(t *testing.T) {
		router, db, _ := setUpAppAndDb()
		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe",
			Age:       13,
		}

		jsonValue, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("Error marshaling JSON: %v", err)
		}
		req, _ := http.NewRequest("POST", "/v1/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var errorMessage errormessage.ErrorMessage
		err = json.Unmarshal(w.Body.Bytes(), &errorMessage)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, errorMessage.ErrorStatusCode)
		assert.Equal(t, "User does not meet minimum age requirement, User email must be properly formatted", errorMessage.ErrorMessage)

		t.Cleanup(func() {
			cleanUpDb(db)
		})
	})

	t.Run("Creating a user with an existing name", func(t *testing.T) {
		router, db, userRepository := setUpAppAndDb()
		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@test.com",
			Age:       25,
		}

		userRepository.CreateUser(&user)

		jsonValue, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("Error marshaling JSON: %v", err)
		}
		req, _ := http.NewRequest("POST", "/v1/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var errorMessage errormessage.ErrorMessage
		err = json.Unmarshal(w.Body.Bytes(), &errorMessage)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusBadRequest, errorMessage.ErrorStatusCode)
		assert.Equal(t, constant.ErrorNameAlreadyExists, errorMessage.ErrorMessage)

		t.Cleanup(func() {
			cleanUpDb(db)
		})
	})

}
