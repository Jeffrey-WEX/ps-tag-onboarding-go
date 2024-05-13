package service

import (
	"net/http"
	"strings"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/errormessage"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
)

type UserService struct {
	userRepository repository.IUserRepository
	userValidation UserValidationService
}

func NewService(userRepository repository.IUserRepository, userValidation UserValidationService) *UserService {
	return &UserService{userRepository, userValidation}
}

func (service UserService) GetUserById(userId string) (*model.User, *errormessage.ErrorMessage) {
	user, err := service.userRepository.GetUserById(userId)

	if err != nil {
		if err.Error() == constant.ErrorUserNotFound {
			errorMessage := errormessage.NewErrorMessage(constant.ErrorUserNotFound, http.StatusNotFound)
			return nil, &errorMessage
		}

		errorMessage := errormessage.NewErrorMessage(constant.ErrorGettingUser, http.StatusInternalServerError)
		return nil, &errorMessage
	}

	return user, nil
}

func (service UserService) CreateUser(user *model.User) (*model.User, *errormessage.ErrorMessage) {
	var errors []string = service.userValidation.ValidateUser(user)
	if len(errors) > 0 {
		errorMessage := errormessage.NewErrorMessage(strings.Join(errors, ", "), http.StatusBadRequest)
		return nil, &errorMessage
	}

	newUser, err := service.userRepository.CreateUser(user)
	if err != nil {
		if err.Error() == constant.ErrorNameAlreadyExists {
			errorMessage := errormessage.NewErrorMessage(constant.ErrorNameAlreadyExists, http.StatusBadRequest)
			return nil, &errorMessage
		}
		errorMessage := errormessage.NewErrorMessage(err.Error(), http.StatusInternalServerError)
		return nil, &errorMessage
	}
	return newUser, nil
}
