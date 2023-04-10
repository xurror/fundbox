package main

import (
	"fmt"

	"getting-to-go/app/config"
	"getting-to-go/app/controllers"
	"getting-to-go/app/models"
	"getting-to-go/app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, world!")

	config, err := config.LoadConfig("app/config/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	err = models.Connect(config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Name)
	if err != nil {
		fmt.Println(err)
	}

	models.RunMigrations()

	router := gin.Default()

	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	userController.Register(router)

	router.Run(fmt.Sprintf(":%s", config.Server.Port))
}
