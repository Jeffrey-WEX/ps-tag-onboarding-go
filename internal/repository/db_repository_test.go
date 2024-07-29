package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository/database/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupDatabase() *mocks.IMongoCollection {
	return new(mocks.IMongoCollection)
}

func TestDbRepository_GetUserByID(t *testing.T) {
	t.Run("Get user successfully", func(t *testing.T) {

		// Arrange
		collectionMock := setupDatabase()
		dbRepo := DbRepository{collection: collectionMock}

		user := model.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@gmail.com",
			Age:       24,
		}

		query := bson.M{"_id": bson.M{"$eq": user.ID}}
		collectionMock.On("FindOne", context.Background(), query).Return(mongo.NewSingleResultFromDocument(&user, nil, nil))

		// Act
		result, _ := dbRepo.GetUserById(user.ID)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "FindOne", context.Background(), query))
		assert.Equal(t, user.ID, result.ID)
		assert.Equal(t, user.FirstName, result.FirstName)
		assert.Equal(t, user.LastName, result.LastName)
		assert.Equal(t, user.Email, result.Email)
		assert.Equal(t, user.Age, result.Age)

	})

	t.Run("Get non-existing user", func(t *testing.T) {
		// Arrange
		collectionMock := setupDatabase()
		dbRepo := DbRepository{collection: collectionMock}

		user := model.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@gmail.com",
			Age:       24,
		}

		query := bson.M{"_id": bson.M{"$eq": user.ID}}
		collectionMock.On("FindOne", context.Background(), query).Return(mongo.NewSingleResultFromDocument(&mongo.SingleResult{}, mongo.ErrNoDocuments, nil))
		expectedError := fmt.Errorf("%s: %v", constant.ErrorUserNotFound, mongo.ErrNoDocuments)
		// Act
		result, err := dbRepo.GetUserById(user.ID)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "FindOne", context.Background(), query))
		assert.EqualError(t, err, expectedError.Error())
		assert.Nil(t, result)
	})

	t.Run("Get user failed with error", func(t *testing.T) {
		// Arrange
		collectionMock := setupDatabase()
		dbRepo := DbRepository{collection: collectionMock}

		user := model.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@gmail.com",
			Age:       24,
		}

		query := bson.M{"_id": bson.M{"$eq": user.ID}}
		collectionMock.On("FindOne", context.Background(), query).Return(mongo.NewSingleResultFromDocument(nil, nil, nil))
		expectedError := fmt.Errorf("%s: %v", constant.ErrorGettingUser, mongo.ErrNilDocument)

		// Act
		result, err := dbRepo.GetUserById(user.ID)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "FindOne", context.Background(), query))
		assert.EqualError(t, err, expectedError.Error())
		assert.Nil(t, result)
	})
}

func TestDbRepository_CreateUser(t *testing.T) {
	t.Run("Create user successfully", func(t *testing.T) {
		// Arrange
		collectionMock := setupDatabase()
		dbRepo := DbRepository{collection: collectionMock}

		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@gmail.com",
			Age:       24,
		}

		query := bson.M{"firstName": bson.M{"$eq": user.FirstName}, "lastName": bson.M{"$eq": user.LastName}}
		collectionMock.On("Find", context.Background(), query).Return(mongo.NewCursorFromDocuments(nil, nil, nil))
		collectionMock.On("InsertOne", context.Background(), &user).Return(&mongo.InsertOneResult{}, nil, nil)

		// Act
		result, err := dbRepo.CreateUser(&user)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "Find", context.Background(), query))
		assert.True(t, collectionMock.AssertCalled(t, "InsertOne", context.Background(), &user))
		assert.Nil(t, err)
		assert.Equal(t, user.FirstName, result.FirstName)
		assert.Equal(t, user.LastName, result.LastName)
		assert.Equal(t, user.Email, result.Email)
		assert.Equal(t, user.Age, result.Age)
	})

	t.Run("Create user failed with error", func(t *testing.T) {
		// Arrange
		collectionMock := setupDatabase()
		dbRepo := DbRepository{collection: collectionMock}

		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@gmail.com",
			Age:       24,
		}

		query := bson.M{"firstName": bson.M{"$eq": user.FirstName}, "lastName": bson.M{"$eq": user.LastName}}
		collectionMock.On("Find", context.Background(), query).Return(mongo.NewCursorFromDocuments(nil, nil, nil))
		collectionMock.On("InsertOne", context.Background(), &user).Return(&mongo.InsertOneResult{}, mongo.ErrClientDisconnected, nil)
		expectedError := fmt.Errorf("%s: %v", constant.ErrorCreatingUser, mongo.ErrClientDisconnected)

		// Act
		result, err := dbRepo.CreateUser(&user)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "Find", context.Background(), query))
		assert.True(t, collectionMock.AssertCalled(t, "InsertOne", context.Background(), &user))
		assert.EqualError(t, err, expectedError.Error())
		assert.Nil(t, result)
	})
}

func TestDbRepository_FindUserByFirstLastName(t *testing.T) {

	t.Run("Find user by first and last name successfully", func(t *testing.T) {
		// Arrange
		collectionMock := setupDatabase()
		dbRepo := DbRepository{collection: collectionMock}

		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@gmail.com",
			Age:       24,
		}

		query := bson.M{"firstName": bson.M{"$eq": user.FirstName}, "lastName": bson.M{"$eq": user.LastName}}
		collectionMock.On("Find", context.Background(), query).Return(mongo.NewCursorFromDocuments([]interface{}{user}, nil, nil))

		// Act
		result, _ := dbRepo.FindUserByFirstLastName(user.FirstName, user.LastName)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "Find", context.Background(), query))
		assert.Equal(t, user.FirstName, result.FirstName)
		assert.Equal(t, user.LastName, result.LastName)
		assert.Equal(t, user.Email, result.Email)
		assert.Equal(t, user.Age, result.Age)
	})

	t.Run("Find user by first and last name failed with error", func(t *testing.T) {
		// Arrange
		collectionMock := setupDatabase()
		dbRepo := DbRepository{collection: collectionMock}

		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "JohnDoe@gmail.com",
			Age:       24,
		}

		query := bson.M{"firstName": bson.M{"$eq": user.FirstName}, "lastName": bson.M{"$eq": user.LastName}}
		collectionMock.On("Find", context.Background(), query).Return(&mongo.Cursor{}, mongo.ErrClientDisconnected)
		expectedError := fmt.Errorf("%s: %v", constant.ErrorFindingUser, mongo.ErrClientDisconnected)

		// Act
		result, err := dbRepo.FindUserByFirstLastName(user.FirstName, user.LastName)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "Find", context.Background(), query))
		assert.EqualError(t, err, expectedError.Error())
		assert.Equal(t, model.User{}, result)
	})
}
