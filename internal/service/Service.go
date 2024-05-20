package service

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/errormessage"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
)

type IService interface {
	GetUserById(userId string) (*model.User, *errormessage.ErrorMessage)
	CreateUser(newUser *model.User) (*model.User, *errormessage.ErrorMessage)
}
