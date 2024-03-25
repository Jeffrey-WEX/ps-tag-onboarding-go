package route

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	controller controller.IController
}

func NewRouter(controller controller.IController) Routes {
	return Routes{controller}
}

func (r *Routes) InitializeRouter(router *gin.Engine) {
	router.GET("/users", r.controller.GetAllUsers)
	router.GET("/users/:id", r.controller.GetUserById)
	router.POST("/users", r.controller.CreateUser)
}
