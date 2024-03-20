package service

import (
	"net/mail"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
)

type UserValidationService struct {
	userRepository repository.UserRepository
}

func NewUserValidationService(userRepository repository.UserRepository) UserValidationService {
	return UserValidationService{userRepository}
}

func (service UserValidationService) ValidateUser(user model.User) (bool, model.User) {
	validateAge(&user)
	validateEmail(&user)
	validateName(&user, service)

	if len(user.ValidationErrors) > 0 {
		return false, user

	}

	return true, user
}

func validateAge(user *model.User) {
	if user.Age < 18 {
		user.ValidationErrors = append(user.ValidationErrors, "User does not meet minimum age requirement")
	}
}

func validateEmail(user *model.User) {
	if user.Email == "" {
		user.ValidationErrors = append(user.ValidationErrors, "User email required")

	} else {
		_, err := mail.ParseAddress(user.Email)

		if err != nil {
			user.ValidationErrors = append(user.ValidationErrors, "User email must be properly formatted")
		}
	}
}

func validateName(user *model.User, service UserValidationService) {
	if user.FirstName == "" || user.LastName == "" {
		user.ValidationErrors = append(user.ValidationErrors, "User first/last names required")
	} else {
		var existingUser = service.userRepository.FindUserByFirstLastName(user.FirstName, user.LastName)

		if existingUser.ID != "" {
			user.ValidationErrors = append(user.ValidationErrors, "User with the same first and last name already exists")
		}
	}
}
