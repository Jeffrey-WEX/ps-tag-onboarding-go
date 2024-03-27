package repository

import (
	"context"
	"errors"
	"log"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
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
		log.Println("User not found: ", err)
		return nil, errors.New(constant.ErrorUserNotFound)
	}

	if err != nil {
		log.Println("Error getting user: ", err)
		return nil, errors.New(constant.ErrorGettingUser)
	}

	return &user, nil
}

func (repo DbRepository) CreateUser(newUser *model.User) (*model.User, error) {
	existingUser, err := repo.FindUserByFirstLastName(newUser.FirstName, newUser.LastName)
	if err != nil {
		return nil, err
	}

	if existingUser.ID != "" {
		return nil, errors.New(constant.ErrorNameAlreadyExists)
	}

	newUser.ID = uuid.New().String()
	_, err = repo.db.Collection("user").InsertOne(context.TODO(), newUser)

	if err != nil {
		return nil, errors.New(constant.ErrorCreatingUser)
	}

	return newUser, nil
}

func (repo DbRepository) FindUserByFirstLastName(firstName string, lastName string) (model.User, error) {
	query := bson.M{"firstName": bson.M{"$eq": firstName}, "lastName": bson.M{"$eq": lastName}}

	cursor, err := repo.db.Collection("user").Find(context.Background(), query)

	if err != nil {
		log.Println("Error finding user: ", err)
		return model.User{}, errors.New(constant.ErrorFindingUser)
	}

	users, err := retrieveUsersFromCursor(cursor)

	if len(users) == 0 {
		return model.User{}, err
	}

	return users[0], nil
}

func retrieveUsersFromCursor(cursor *mongo.Cursor) ([]model.User, error) {
	var users []model.User

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user model.User
		err := cursor.Decode(&user)

		if err != nil {
			log.Println("Error decoding user: ", err)
			return nil, errors.New(constant.ErrorDecodingUser)
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error retrieving users: ", err)
		return nil, errors.New(constant.ErrorRetrievingUsers)
	}

	return users, nil
}
