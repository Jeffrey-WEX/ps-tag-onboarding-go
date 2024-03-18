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

type DbRepository struct {
}

func NewRepository() DbRepository {
	return DbRepository{}
}

func (repo DbRepository) GetUserById(id string) (*model.User, error) {

	for i, user := range users {
		if user.ID == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

func (repo DbRepository) GetAllUsers() []model.User {
	return users
}

func (repo DbRepository) AddUser(newUser model.User) model.User {
	newUser.ID = uuid.New().String()
	users = append(users, newUser)
	return newUser
}

func (repo DbRepository) FindUserByFirstLastName(firstName string, lastName string) model.User {
	for i, user := range users {
		if user.FirstName == firstName && user.LastName == lastName {
			return users[i]
		}
	}

	return model.User{}
}
