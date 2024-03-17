package controller

import (
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service"
)

type Controller struct {
	service service.Service
}

// func GetAllUsers(context *gin.Context) {
// 	context.IndentedJSON(http.StatusOk, users)
// }

// func getUserById(context *gin.Context) {
// 	context
// }
