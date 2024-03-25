package service

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
)

type UserService struct {
	userRepository repository.IUserRepository
	UserValidation UserValidationService
}

func NewService(userRepository repository.IUserRepository, userValidation UserValidationService) UserService {
	return UserService{userRepository, userValidation}
}

func (service UserService) GetAllUsers() []model.User {
	return service.userRepository.GetAllUsers()
}

func (service UserService) GetUserById(userId string) (*model.User, error) {
	return service.userRepository.GetUserById(userId)
}

func (service UserService) CreateUser(newUser model.User) model.User {
	valid, user := service.UserValidation.ValidateUser(newUser)

	if !valid {
		return user
	} else {
		return service.userRepository.CreateUser(newUser)
	}
}
