package repository

import (
	"testing"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func setUpDatabase() *mongo.Database {
	return &database.MongoCollectionMock{}
}

func TestDbRepository_GetUserByID(t *testing.T) {
	t.Run("Get user successfully", func(t *testing.T) {
		// Arrange
		db := setUpDatabase()
		dbRepo := NewRepository(db)
		user := &model.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "",
			Age:       25,
		}

	})
}
