package repository

import "github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"

type IUserRepository interface {
	GetUserById(userId string) (*model.User, error)
	CreateUser(newUser *model.User) (*model.User, error)
	FindUserByFirstLastName(firstName string, lastName string) (model.User, error)
}
