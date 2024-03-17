package service

import (
	"net/http"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
	"github.com/gin-gonic/gin"
)

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return Service{repository}
}

func (service Service) GetAllUsers(context *gin.Context) {
	users := service.repository.GetAllUsers()
	context.IndentedJSON(http.StatusOK, users)
}

func (service Service) GetUserById(context *gin.Context) {
	id := context.Param("id")
	user, err := service.repository.GetUserById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}

	context.IndentedJSON(http.StatusOK, user)
}

func (service Service) AddUser(context *gin.Context) {
	var newUser model.User

	if err := context.BindJSON(&newUser); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid user object"})
		return
	}

	service.repository.AddUser(newUser)
	context.IndentedJSON(http.StatusCreated, newUser)
}
