package main

import (
	"fmt"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Application!")
	var repository = repository.NewRepository()
	var service = service.NewService(repository)
	startHttpServer(service)
}

func startHttpServer(service service.Service) {
	router := gin.Default()
	router.GET("/users", service.GetAllUsers)
	router.GET("/users/:id", service.GetUserById)
	router.POST("/users", service.AddUser)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
