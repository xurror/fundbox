package main

import (
	"fmt"

	"getting-to-go/config"
	"getting-to-go/controllers"
	"getting-to-go/models"
	"getting-to-go/services"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, world!")

	config, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	err = models.Connect(config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Name)
	if err != nil {
		fmt.Println(err)
	}

	models.RunMigrations()

	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	fundService := services.NewFundService()
	fundController := controllers.NewFundController(fundService)

	contributionService := services.NewContributionService()
	contributionController := controllers.NewContributionController(contributionService)

	router := gin.Default()

	api := router.Group("/api")
	apiV1 := api.Group("/v1")

	userController.Register(apiV1)
	fundController.Register(apiV1)
	contributionController.Register(apiV1)

	router.Run(fmt.Sprintf(":%s", config.Server.Port))
}
