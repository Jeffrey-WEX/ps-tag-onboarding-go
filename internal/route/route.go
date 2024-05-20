package route

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller"
	"github.com/gin-gonic/gin"
)

type IRoutes interface {
	InitializeRouter(router *gin.Engine)
}

type Routes struct {
	controller controller.IController
}

func NewRoutes(controller controller.IController) Routes {
	return Routes{controller}
}

func (r *Routes) InitializeRouter(router *gin.Engine) {
	router.GET("/v1/users/:id", r.controller.GetUserById)
	router.POST("/v1/users", r.controller.CreateUser)
}
