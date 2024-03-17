package repository

import (
	"errors"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
)

var users = []model.User{
	{ID: "1", FirstName: "John", LastName: "Doe", Email: "John.Doe@gmail.com", Age: 40},
	{ID: "2", FirstName: "Matt", LastName: "White", Email: "Matt.White@gmail.com", Age: 21},
	{ID: "3", FirstName: "Connor", LastName: "Pan", Email: "Connor.Pan@gmail.com", Age: 35},
}

type Repository struct {
	repo *Repository
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) GetUserById(id string) (*model.User, error) {

	for i, user := range users {
		if user.ID == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("User not found")
}

func (r Repository) GetAllUsers() []model.User {
	return users
}

func (r Repository) AddUser(newUser model.User) {
	users = append(users, newUser)
}
