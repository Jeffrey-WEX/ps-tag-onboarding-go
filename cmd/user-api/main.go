package main

import (
	"fmt"

	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/repository/database"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/route"
	"github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Application!")
	var userRepository = repository.NewRepository(database.NewDatabase())
	var userValidator = service.NewUserValidationService(userRepository)
	var userService = service.NewService(userRepository, userValidator)
	var userController = controller.NewController(userService)
	var routes = route.NewRoutes(userController)
	startHttpServer(routes)
}

func startHttpServer(routes route.Routes) {
	router := gin.Default()
	routes.InitializeRouter(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
