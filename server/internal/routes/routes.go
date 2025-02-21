package routes

import (
	"community-funds/internal/config"
	"community-funds/internal/handlers"
	"community-funds/internal/middlewares"
	"community-funds/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	_ "community-funds/docs" // Import Swagger docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Cfg                 *config.Config
	Logger              *logrus.Logger
	UserRepo            *repositories.UserRepository
	FundHandler         *handlers.FundHandler
	ContributionHandler *handlers.ContributionHandler
}

func NewRouter(
	cfg *config.Config,
	logger *logrus.Logger,
	userRepo *repositories.UserRepository,
	fundHandler *handlers.FundHandler,
	contributionHandler *handlers.ContributionHandler,
) *Router {
	return &Router{
		Cfg:                 cfg,
		Logger:              logger,
		UserRepo:            userRepo,
		FundHandler:         fundHandler,
		ContributionHandler: contributionHandler,
	}
}

func (r *Router) SetupRoutes(router *gin.Engine) {

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	api.Use(middlewares.AuthMiddleware(r.UserRepo, r.Cfg, r.Logger))
	{
		funds := api.Group("/funds")
		{
			funds.POST("", r.FundHandler.CreateFund)
			funds.GET("", r.FundHandler.GetFunds)
			funds.GET("/contributed", r.FundHandler.GetContributedFunds)
		}

		contributions := api.Group("/contributions")
		{
			contributions.POST("", r.ContributionHandler.CreateContribution)
			contributions.GET("", r.ContributionHandler.GetContributions)
		}
	}
}
