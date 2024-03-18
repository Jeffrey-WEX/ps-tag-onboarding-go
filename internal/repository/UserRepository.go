package repository

import (
	"errors"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/google/uuid"
)

var users = []model.User{
	{ID: uuid.New().String(), FirstName: "John", LastName: "Doe", Email: "John.Doe@gmail.com", Age: 40},
	{ID: uuid.New().String(), FirstName: "Matt", LastName: "White", Email: "Matt.White@gmail.com", Age: 21},
	{ID: uuid.New().String(), FirstName: "Connor", LastName: "Pan", Email: "Connor.Pan@gmail.com", Age: 35},
}

type UserRepository struct {
	// TODO: change this to interface
	// TODO: create a new repository class for mongoDB
}

func NewRepository() UserRepository {
	return UserRepository{}
}

func (r UserRepository) GetUserById(id string) (*model.User, error) {

	for i, user := range users {
		if user.ID == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

func (r UserRepository) GetAllUsers() []model.User {
	return users
}

func (r UserRepository) AddUser(newUser model.User) model.User {
	newUser.ID = uuid.New().String()
	users = append(users, newUser)
	return newUser
}

func (r UserRepository) FindUserByFirstLastName(firstName string, lastName string) model.User {
	for i, user := range users {
		if user.FirstName == firstName && user.LastName == lastName {
			return users[i]
		}
	}

	return model.User{}
}
