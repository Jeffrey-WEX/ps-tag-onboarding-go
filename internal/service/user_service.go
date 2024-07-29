package service

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/errormessage"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
)

var (
	once    sync.Once
	service *UserService
)

type IUserService interface {
	GetUserById(userId string) (*model.User, *errormessage.ErrorMessage)
	CreateUser(newUser *model.User) (*model.User, *errormessage.ErrorMessage)
}

type UserService struct {
	dbRepository   repository.IDbRepository
	userValidation UserValidationService
}

func NewService(dbRepository repository.IDbRepository, userValidation UserValidationService) *UserService {
	once.Do(func() {
		service = &UserService{dbRepository, userValidation}
	})
	return service
}

func (service UserService) GetUserById(userId string) (*model.User, *errormessage.ErrorMessage) {
	user, err := service.dbRepository.GetUserById(userId)

	if err != nil {
		if err.Error() == constant.ErrorUserNotFound {
			log.Println("Failed to get user, error: ", err)
			errorMessage := errormessage.NewErrorMessage(constant.ErrorUserNotFound, http.StatusNotFound)
			return nil, &errorMessage
		}

		log.Println("Failed to get user, error: ", err)
		errorMessage := errormessage.NewErrorMessage(constant.ErrorGettingUser, http.StatusInternalServerError)
		return nil, &errorMessage
	}

	return user, nil
}

func (service UserService) CreateUser(user *model.User) (*model.User, *errormessage.ErrorMessage) {
	var errors []string = service.userValidation.ValidateUser(user)
	if len(errors) > 0 {
		log.Printf("Validation failed for user creation: %v", strings.Join(errors, ", "))
		errorMessage := errormessage.NewErrorMessage(strings.Join(errors, ", "), http.StatusBadRequest)
		return nil, &errorMessage
	}

	newUser, err := service.dbRepository.CreateUser(user)
	if err != nil {
		if err.Error() == constant.ErrorNameAlreadyExists {
			log.Println("Failed to create user, error: ", err)
			errorMessage := errormessage.NewErrorMessage(constant.ErrorNameAlreadyExists, http.StatusBadRequest)
			return nil, &errorMessage
		}

		log.Println("Failed to create user, error: ", err)
		errorMessage := errormessage.NewErrorMessage(err.Error(), http.StatusInternalServerError)
		return nil, &errorMessage
	}
	return newUser, nil
}
