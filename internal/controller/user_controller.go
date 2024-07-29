package controller

import (
	"log"
	"net/http"
	"sync"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/constant"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	once       sync.Once
	controller *UserController
)

type IUserController interface {
	GetUserById(context *gin.Context)
	CreateUser(context *gin.Context)
}

type UserController struct {
	userService service.IUserService
}

func NewController(userService service.IUserService) *UserController {
	once.Do(func() {
		controller = &UserController{userService}
	})
	return controller
}

func (controller UserController) GetUserById(context *gin.Context) {
	id := context.Param("id")
	user, errorMessage := controller.userService.GetUserById(id)

	if errorMessage != nil {
		log.Println("Failed to get user, error: ", errorMessage.ErrorMessage)
		context.IndentedJSON(errorMessage.ErrorStatusCode, gin.H{"status_code": errorMessage.ErrorStatusCode, "message": errorMessage.ErrorMessage})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func (controller UserController) CreateUser(context *gin.Context) {
	var user model.User

	if err := context.BindJSON(&user); err != nil {
		log.Println("Failed to bind JSON, error: ", err)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "message": constant.ErrorInvalidUserObject})
		return
	}

	newUser, errorMessage := controller.userService.CreateUser(&user)

	if errorMessage != nil {
		log.Println("Failed to create user, error: ", errorMessage.ErrorMessage)
		context.IndentedJSON(errorMessage.ErrorStatusCode, gin.H{"status_code": errorMessage.ErrorStatusCode, "message": errorMessage.ErrorMessage})
		return
	} else {
		context.IndentedJSON(http.StatusCreated, newUser)
	}

}
