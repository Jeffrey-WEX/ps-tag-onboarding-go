package controller

import "github.com/gin-gonic/gin"

type IController interface {
	GetUserById(context *gin.Context)
	CreateUser(context *gin.Context)
}
