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

func TestDbRepository_GetUserById(t *testing.T) {
	user := model.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@gmail.com",
		Age:       24,
	}

	// Test Get user successfully
	mockFindOneFuncReturnUser := func(collectionMock *mocks.IMongoCollection) {
		query := bson.M{"_id": bson.M{"$eq": user.ID}}
		collectionMock.On("FindOne", context.Background(), query).Return(mongo.NewSingleResultFromDocument(user, nil, nil))
	}

	// Test Get non-existing user
	mockFindOneFuncReturnEmpty := func(collectionMock *mocks.IMongoCollection) {
		query := bson.M{"_id": bson.M{"$eq": user.ID}}
		collectionMock.On("FindOne", context.Background(), query).Return(mongo.NewSingleResultFromDocument(&mongo.SingleResult{}, mongo.ErrNoDocuments, nil))
	}

	// Test Get user failed with error
	mockFindOneFuncReturnError := func(collectionMock *mocks.IMongoCollection) {
		query := bson.M{"_id": bson.M{"$eq": user.ID}}
		collectionMock.On("FindOne", context.Background(), query).Return(mongo.NewSingleResultFromDocument(nil, nil, nil))
	}

	testCases := []struct {
		name            string
		userID          string
		mockFindOneFunc func(*mocks.IMongoCollection)
		expectedUser    *model.User
		expectedError   error
	}{
		{
			name:            "Get user successfully",
			userID:          user.ID,
			mockFindOneFunc: mockFindOneFuncReturnUser,
			expectedUser:    &user,
			expectedError:   nil,
		},
		{
			name:            "Get non-existing user",
			userID:          user.ID,
			mockFindOneFunc: mockFindOneFuncReturnEmpty,
			expectedUser:    nil,
			expectedError:   fmt.Errorf("%s: %v", constant.ErrorUserNotFound, mongo.ErrNoDocuments),
		},
		{
			name:            "Get user failed with error",
			userID:          user.ID,
			mockFindOneFunc: mockFindOneFuncReturnError,
			expectedUser:    nil,
			expectedError:   fmt.Errorf("%s: %v", constant.ErrorGettingUser, mongo.ErrNilDocument),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			collectionMock := setupDatabase()
			dbRepo := DbRepository{collection: collectionMock}
			tc.mockFindOneFunc(collectionMock)

			// Act
			result, err := dbRepo.GetUserById(tc.userID)

			// Assert
			assert.True(t, collectionMock.AssertCalled(t, "FindOne", context.Background(), bson.M{"_id": bson.M{"$eq": tc.userID}}))
			assert.Equal(t, tc.expectedUser, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDbRepository_CreateUser(t *testing.T) {
	user := model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@gmail.com",
		Age:       24,
	}

	mockFindOneFuncReturnEmpty := func(collectionMock *mocks.IMongoCollection) {
		query := bson.M{"firstName": bson.M{"$eq": user.FirstName}, "lastName": bson.M{"$eq": user.LastName}}
		collectionMock.On("Find", context.Background(), query).Return(mongo.NewCursorFromDocuments(nil, nil, nil))
	}

	// Test Create user successfully
	mockInsertOneFuncReturnSuccess := func(collectionMock *mocks.IMongoCollection) {
		collectionMock.On("InsertOne", context.Background(), &user).Return(&mongo.InsertOneResult{}, nil)
	}

	// Test Create user failed with error
	mockInsertOneFuncReturnError := func(collectionMock *mocks.IMongoCollection) {
		collectionMock.On("InsertOne", context.Background(), &user).Return(&mongo.InsertOneResult{}, mongo.ErrClientDisconnected)
	}

	testCases := []struct {
		name              string
		user              *model.User
		mockFindOneFunc   func(*mocks.IMongoCollection)
		mockInsertOneFunc func(*mocks.IMongoCollection)
		expectedUser      *model.User
		expectedError     error
	}{
		{
			name:              "Create user successfully",
			user:              &user,
			mockFindOneFunc:   mockFindOneFuncReturnEmpty,
			mockInsertOneFunc: mockInsertOneFuncReturnSuccess,
			expectedUser:      &user,
			expectedError:     nil,
		},
		{
			name:              "Create user failed with error",
			user:              &user,
			mockFindOneFunc:   mockFindOneFuncReturnEmpty,
			mockInsertOneFunc: mockInsertOneFuncReturnError,
			expectedUser:      nil,
			expectedError:     fmt.Errorf("%s: %v", constant.ErrorCreatingUser, mongo.ErrClientDisconnected),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			collectionMock := setupDatabase()
			dbRepo := DbRepository{collection: collectionMock}
			tc.mockFindOneFunc(collectionMock)
			tc.mockInsertOneFunc(collectionMock)

			// Act
			result, err := dbRepo.CreateUser(tc.user)

			// Assert
			query := bson.M{"firstName": bson.M{"$eq": tc.user.FirstName}, "lastName": bson.M{"$eq": tc.user.LastName}}
			assert.True(t, collectionMock.AssertCalled(t, "Find", context.Background(), query))
			assert.True(t, collectionMock.AssertCalled(t, "InsertOne", context.Background(), tc.user))
			assert.Equal(t, tc.expectedUser, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDbRepository_FindUserByFirstLastName(t *testing.T) {
	user := model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@gmail.com",
		Age:       24,
	}

	// Test Find user by first and last name successfully
	mockFindFuncReturnUser := func(collectionMock *mocks.IMongoCollection) {
		query := bson.M{"firstName": bson.M{"$eq": user.FirstName}, "lastName": bson.M{"$eq": user.LastName}}
		collectionMock.On("Find", context.Background(), query).Return(mongo.NewCursorFromDocuments([]interface{}{user}, nil, nil))
	}

	// Test Find user by first and last name failed with error
	mockFindFuncReturnError := func(collectionMock *mocks.IMongoCollection) {
		query := bson.M{"firstName": bson.M{"$eq": user.FirstName}, "lastName": bson.M{"$eq": user.LastName}}
		collectionMock.On("Find", context.Background(), query).Return(&mongo.Cursor{}, mongo.ErrClientDisconnected)
	}

	testCases := []struct {
		name          string
		firstName     string
		lastName      string
		mockFindFunc  func(*mocks.IMongoCollection)
		expectedUser  model.User
		expectedError error
	}{
		{
			name:          "Find user by first and last name successfully",
			firstName:     user.FirstName,
			lastName:      user.LastName,
			mockFindFunc:  mockFindFuncReturnUser,
			expectedUser:  user,
			expectedError: nil,
		},
		{
			name:          "Find user by first and last name failed with error",
			firstName:     user.FirstName,
			lastName:      user.LastName,
			mockFindFunc:  mockFindFuncReturnError,
			expectedUser:  model.User{},
			expectedError: fmt.Errorf("%s: %v", constant.ErrorFindingUser, mongo.ErrClientDisconnected),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			collectionMock := setupDatabase()
			dbRepo := DbRepository{collection: collectionMock}
			tc.mockFindFunc(collectionMock)

			// Act
			result, err := dbRepo.FindUserByFirstLastName(tc.firstName, tc.lastName)

			// Assert
			query := bson.M{"firstName": bson.M{"$eq": tc.firstName}, "lastName": bson.M{"$eq": tc.lastName}}
			assert.True(t, collectionMock.AssertCalled(t, "Find", context.Background(), query))
			assert.Equal(t, tc.expectedUser, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
