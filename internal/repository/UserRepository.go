package repository

import "github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"

type UserRepository interface {
	GetAllUsers() []model.User
	GetUserById(userId string) (*model.User, error)
	AddUser(newUser model.User) model.User
}
