package controller

import (
	"net/http"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/model"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewController(userService service.UserService) UserController {
	return UserController{userService}
}

func (controller UserController) GetAllUsers(context *gin.Context) {
	users := controller.userService.GetAllUsers()
	context.IndentedJSON(http.StatusOK, users)
}

func (controller UserController) GetUserById(context *gin.Context) {
	id := context.Param("id")
	user, err := controller.userService.GetUserById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}

	context.IndentedJSON(http.StatusOK, user)
}

func (controller UserController) AddUser(context *gin.Context) {
	var newUser model.User

	if err := context.BindJSON(&newUser); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid user object"})
		return
	}

	newUser = controller.userService.AddUser(newUser)

	if len(newUser.ValidationErrors) > 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"errors": &newUser.ValidationErrors})
	} else {
		context.IndentedJSON(http.StatusCreated, newUser)
	}

}
