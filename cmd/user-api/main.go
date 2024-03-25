package main

import (
	"fmt"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository/database"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Application!")
	var userRepository = repository.NewRepository(database.NewDatabase())
	var userValidator = service.NewUserValidationService(userRepository)
	var userService = service.NewService(userRepository, userValidator)
	var userController = controller.NewController(userService)
	startHttpServer(userController)
}

func startHttpServer(userController controller.UserController) {
	router := gin.Default()
	router.GET("/users", userController.GetAllUsers)
	router.GET("/users/:id", userController.GetUserById)
	router.POST("/users", userController.CreateUser)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
