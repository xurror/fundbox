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
	cfg                 *config.Config
	logger              *logrus.Logger
	actuatorHandler     *handlers.ActuatorHandler
	userRepo            *repositories.UserRepository
	fundHandler         *handlers.FundHandler
	contributionHandler *handlers.ContributionHandler
}

func NewRouter(
	cfg *config.Config,
	logger *logrus.Logger,
	actuatorHandler *handlers.ActuatorHandler,
	contributionHandler *handlers.ContributionHandler,
	fundHandler *handlers.FundHandler,
	userRepo *repositories.UserRepository,
) *Router {
	return &Router{
		cfg:                 cfg,
		logger:              logger,
		actuatorHandler:     actuatorHandler,
		contributionHandler: contributionHandler,
		fundHandler:         fundHandler,
		userRepo:            userRepo,
	}
}

func (r *Router) SetupRoutes(router *gin.Engine) {

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Actuator Endpoints
	actuator := router.Group("/actuator")
	{
		actuator.GET("/health", r.actuatorHandler.HealthCheck)
		actuator.GET("/ready", r.actuatorHandler.ReadinessCheck)
		// actuator.POST("/shutdown", r.actuatorHandler.Shutdown)
		// actuator.POST("/restart", r.actuatorHandler.Restart)
	}

	api := router.Group("/api")
	api.Use(middlewares.AuthMiddleware(r.userRepo, r.cfg, r.logger))
	{
		funds := api.Group("/funds")
		{
			funds.GET("", r.fundHandler.GetFunds)
			funds.POST("", r.fundHandler.CreateFund)
			funds.GET("/:fundId", r.fundHandler.GetFund)
			funds.GET("/contributed", r.fundHandler.GetContributedFunds)
		}

		contributions := api.Group("/contributions")
		{
			contributions.GET("", r.contributionHandler.GetContributions)
			contributions.POST("", r.contributionHandler.CreateContribution)
		}
	}
}
