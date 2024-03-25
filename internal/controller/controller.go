package controller

import "github.com/gin-gonic/gin"

type IController interface {
	GetAllUsers(context *gin.Context)
	GetUserById(context *gin.Context)
	CreateUser(context *gin.Context)
}
