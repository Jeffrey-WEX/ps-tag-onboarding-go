package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository/database/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDatabase() *mocks.IMongoCollectionMock {
	return &mocks.IMongoCollectionMock{}
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
		collectionMock.On("FindOne", context.TODO(), query, []*options.FindOneOptions(nil)).Return(mongo.NewSingleResultFromDocument(&user, nil, nil))

		// Act
		result, _ := dbRepo.GetUserById(user.ID)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "FindOne", context.TODO(), query, []*options.FindOneOptions(nil)))
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
		collectionMock.On("FindOne", context.TODO(), query, []*options.FindOneOptions(nil)).Return(mongo.NewSingleResultFromDocument(&mongo.SingleResult{}, mongo.ErrNoDocuments, nil))
		// Act
		result, err := dbRepo.GetUserById(user.ID)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "FindOne", context.TODO(), query, []*options.FindOneOptions(nil)))
		assert.EqualError(t, err, constant.ErrorUserNotFound)
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
		collectionMock.On("FindOne", context.TODO(), query, []*options.FindOneOptions(nil)).Return(mongo.NewSingleResultFromDocument(nil, nil, nil))

		// Act
		result, err := dbRepo.GetUserById(user.ID)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "FindOne", context.TODO(), query, []*options.FindOneOptions(nil)))
		assert.EqualError(t, err, constant.ErrorGettingUser)
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
		collectionMock.On("Find", context.Background(), query, []*options.FindOptions(nil)).Return(mongo.NewCursorFromDocuments(nil, nil, nil))
		collectionMock.On("InsertOne", context.TODO(), &user, []*options.InsertOneOptions(nil)).Return(&mongo.InsertOneResult{}, nil, nil)

		// Act
		result, err := dbRepo.CreateUser(&user)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "Find", context.Background(), query, []*options.FindOptions(nil)))
		assert.True(t, collectionMock.AssertCalled(t, "InsertOne", context.TODO(), &user, []*options.InsertOneOptions(nil)))
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
		collectionMock.On("Find", context.Background(), query, []*options.FindOptions(nil)).Return(mongo.NewCursorFromDocuments(nil, nil, nil))
		collectionMock.On("InsertOne", context.TODO(), &user, []*options.InsertOneOptions(nil)).Return(&mongo.InsertOneResult{}, errors.New(constant.ErrorCreatingUser), nil)

		// Act
		result, err := dbRepo.CreateUser(&user)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "Find", context.Background(), query, []*options.FindOptions(nil)))
		assert.True(t, collectionMock.AssertCalled(t, "InsertOne", context.TODO(), &user, []*options.InsertOneOptions(nil)))
		assert.EqualError(t, err, constant.ErrorCreatingUser)
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
		collectionMock.On("Find", context.Background(), query, []*options.FindOptions(nil)).Return(mongo.NewCursorFromDocuments([]interface{}{user}, nil, nil))

		// Act
		result, _ := dbRepo.FindUserByFirstLastName(user.FirstName, user.LastName)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "Find", context.Background(), query, []*options.FindOptions(nil)))
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
		collectionMock.On("Find", context.Background(), query, []*options.FindOptions(nil)).Return(&mongo.Cursor{}, errors.New(constant.ErrorFindingUser))

		// Act
		result, err := dbRepo.FindUserByFirstLastName(user.FirstName, user.LastName)

		// Assert
		assert.True(t, collectionMock.AssertCalled(t, "Find", context.Background(), query, []*options.FindOptions(nil)))
		assert.EqualError(t, err, constant.ErrorFindingUser)
		assert.Equal(t, model.User{}, result)
	})
}
