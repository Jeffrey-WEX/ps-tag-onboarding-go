package service

import (
	"net/mail"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
)

type UserValidationService struct {
}

func NewUserValidationService() UserValidationService {
	return UserValidationService{}
}

func (service UserValidationService) ValidateUser(user *model.User) []string {
	var errors []string

	var error string = validateAge(user)
	if error != "" {
		errors = append(errors, error)
	}

	error = validateEmail(user)
	if error != "" {
		errors = append(errors, error)
	}

	error = validateName(user)
	if error != "" {
		errors = append(errors, error)
	}

	return errors
}

func validateAge(user *model.User) string {
	if user.Age < 18 {
		return constant.ErrorAgeMinimum
	}
	return ""
}

func validateEmail(user *model.User) string {
	if user.Email == "" {
		return constant.ErrorEmailRequired
	} else {
		_, err := mail.ParseAddress(user.Email)

		if err != nil {
			return constant.ErrorEmailInvalidFormat
		}
	}

	return ""
}

func validateName(user *model.User) string {
	if user.FirstName == "" || user.LastName == "" {
		return constant.ErrorNameRequired
	}

	return ""
}
