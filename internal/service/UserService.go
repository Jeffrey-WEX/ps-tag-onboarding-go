package service

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) UserService {
	return UserService{userRepository}
}

func (service UserService) GetAllUsers() []model.User {
	return service.userRepository.GetAllUsers()
}

func (service UserService) GetUserById(userId string) (*model.User, error) {
	return service.userRepository.GetUserById(userId)
}

func (service UserService) AddUser(newUser model.User) {
	service.userRepository.AddUser(newUser)
}
