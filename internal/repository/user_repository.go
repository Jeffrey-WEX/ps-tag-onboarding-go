package repository

import "github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"

type IUserRepository interface {
	GetAllUsers() []model.User
	GetUserById(userId string) (*model.User, error)
	CreateUser(newUser model.User) model.User
	FindUserByFirstLastName(firstName string, lastName string) model.User
}
