package route

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	UserController controller.UserController
}

func NewRouter(userController controller.UserController) Routes {
	return Routes{UserController: userController}
}

func (r *Routes) InitializeRouter(router *gin.Engine) {
	router.GET("/users", r.UserController.GetAllUsers)
	router.GET("/users/:id", r.UserController.GetUserById)
	router.POST("/users", r.UserController.CreateUser)
}
