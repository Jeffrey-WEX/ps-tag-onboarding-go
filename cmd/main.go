package main

import (
	"fmt"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Application!")
	var repository = repository.NewRepository()
	var service = service.NewService(repository)
	var controller = controller.NewController(service)
	startHttpServer(controller)
}

func startHttpServer(controller controller.UserController) {
	router := gin.Default()
	router.GET("/users", controller.GetAllUsers)
	router.GET("/users/:id", controller.GetUserById)
	router.POST("/users", controller.AddUser)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
