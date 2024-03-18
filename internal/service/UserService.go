package service

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
)

type UserService struct {
	userRepository repository.UserRepository
	UserValidation UserValidationService
}

func NewService(userRepository repository.UserRepository, userValidation UserValidationService) UserService {
	return UserService{userRepository, userValidation}
}

func (service UserService) GetAllUsers() []model.User {
	return service.userRepository.GetAllUsers()
}

func (service UserService) GetUserById(userId string) (*model.User, error) {
	return service.userRepository.GetUserById(userId)
}

func (service UserService) AddUser(newUser model.User) model.User {
	// TODO: Validate user
	valid, user := service.UserValidation.ValidateUser(newUser)

	if !valid {
		return user
	} else {
		return service.userRepository.AddUser(newUser)
	}
}
