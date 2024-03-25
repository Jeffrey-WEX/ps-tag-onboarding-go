package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbRepository struct {
	db *mongo.Database
}

func NewRepository(database *mongo.Database) DbRepository {
	return DbRepository{database}
}

func (repo DbRepository) GetUserById(id string) (*model.User, error) {
	var user model.User
	query := bson.M{"_id": bson.M{"$eq": id}}
	err := repo.db.Collection("user").FindOne(context.TODO(), query).Decode(&user)

	if err == mongo.ErrNoDocuments {
		fmt.Println("User not found: ", err)
		return nil, errors.New("user not found")
	}

	if err != nil {
		fmt.Println("Error getting user: ", err)
		return nil, errors.New("error getting user")
	}

	return &user, nil
}

func (repo DbRepository) CreateUser(newUser model.User) model.User {
	newUser.ID = uuid.New().String()
	_, err := repo.db.Collection("user").InsertOne(context.TODO(), newUser)

	if err != nil {
		fmt.Println("Error creating user: ", err)
	}

	return newUser
}

func (repo DbRepository) FindUserByFirstLastName(firstName string, lastName string) model.User {
	query := bson.M{"firstName": bson.M{"$eq": firstName}, "lastName": bson.M{"$eq": lastName}}

	cursor, err := repo.db.Collection("user").Find(context.Background(), query)

	if err != nil {
		fmt.Println("Error finding user: ", err)
	}

	var users []model.User = retrieveUsersFromCursor(cursor)

	if len(users) == 0 {
		return model.User{}
	}

	return users[0]
}

func retrieveUsersFromCursor(cursor *mongo.Cursor) []model.User {
	var users []model.User

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user model.User
		err := cursor.Decode(&user)

		if err != nil {
			fmt.Println("Error decoding user: ", err)
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println("Error retrieving users: ", err)
	}

	return users
}
