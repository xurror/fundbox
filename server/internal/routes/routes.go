package routes

import (
	"community-funds/internal/handlers"

	"github.com/gin-gonic/gin"

	_ "community-funds/docs" // Import Swagger docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	FundHandler *handlers.FundHandler
	UserHandler *handlers.UserHandler
}

func NewRouter(fundHandler *handlers.FundHandler, userHandler *handlers.UserHandler) *Router {
	return &Router{FundHandler: fundHandler, UserHandler: userHandler}
}

func (r *Router) SetupRoutes(router *gin.Engine) {

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		api.GET("/funds", r.FundHandler.GetFunds)
		api.GET("/users", r.UserHandler.GetUsers)
	}
}
