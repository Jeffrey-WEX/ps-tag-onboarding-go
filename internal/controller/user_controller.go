package controller

import (
	"net/http"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.IService
}

func NewController(service service.IService) UserController {
	return UserController{service}
}

func (controller UserController) GetUserById(context *gin.Context) {
	id := context.Param("id")
	user, errorMessage := controller.service.GetUserById(id)

	if errorMessage != nil {
		context.IndentedJSON(errorMessage.StatusCode(), gin.H{"message": errorMessage.Message()})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func (controller UserController) CreateUser(context *gin.Context) {
	var user model.User

	if err := context.BindJSON(&user); err != nil {

		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": constant.ErrorInvalidUserObject})
		return
	}

	newUser, errorMessage := controller.service.CreateUser(&user)

	if errorMessage != nil {
		context.IndentedJSON(errorMessage.StatusCode(), gin.H{"errors": errorMessage.Message()})
	} else {
		context.IndentedJSON(http.StatusCreated, newUser)
	}

}
