package routes

import (
	"community-funds/internal/config"
	"community-funds/internal/handlers"
	"community-funds/internal/middlewares"
	"community-funds/internal/repositories"

	"github.com/gin-gonic/gin"

	_ "community-funds/docs" // Import Swagger docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Cfg                 *config.Config
	UserRepo            *repositories.UserRepository
	FundHandler         *handlers.FundHandler
	ContributionHandler *handlers.ContributionHandler
}

func NewRouter(
	cfg *config.Config,
	userRepo *repositories.UserRepository,
	fundHandler *handlers.FundHandler,
	contributionHandler *handlers.ContributionHandler,
) *Router {
	return &Router{
		Cfg:                 cfg,
		UserRepo:            userRepo,
		FundHandler:         fundHandler,
		ContributionHandler: contributionHandler,
	}
}

func (r *Router) SetupRoutes(router *gin.Engine) {

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{

		// Protected routes (Only Fund Managers)
		protected := api.Group("/funds")
		protected.Use(middlewares.AuthMiddleware(r.UserRepo, r.Cfg))
		protected.POST("", r.FundHandler.CreateFund)

		// Contributions (Anonymous allowed)
		api.POST("/contributions", r.ContributionHandler.CreateContribution)
	}
}
