package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"community-funds/api/handlers"
	"community-funds/api/middlewares"
	"community-funds/config"

	// _ "community-funds/docs" // Import Swagger docs
	"community-funds/pkg/repositories"
)

type Router struct {
	cfg                 *config.Config
	log                 *logrus.Logger
	actuatorHandler     *handlers.ActuatorHandler
	userRepo            *repositories.UserRepository
	fundHandler         *handlers.FundHandler
	contributionHandler *handlers.ContributionHandler
}

func NewRouter(
	cfg *config.Config,
	log *logrus.Logger,
	actuatorHandler *handlers.ActuatorHandler,
	contributionHandler *handlers.ContributionHandler,
	fundHandler *handlers.FundHandler,
	userRepo *repositories.UserRepository,
) *Router {
	return &Router{
		cfg:                 cfg,
		log:                 log,
		actuatorHandler:     actuatorHandler,
		contributionHandler: contributionHandler,
		fundHandler:         fundHandler,
		userRepo:            userRepo,
	}
}

func (r *Router) SetupRoutes(router *fiber.App) {
	// Actuator Endpoints
	actuator := router.Group("/actuator")
	{
		actuator.Get("/health", r.actuatorHandler.HealthCheck)
		actuator.Get("/ready", r.actuatorHandler.ReadinessCheck)
		// actuator.POST("/shutdown", r.actuatorHandler.Shutdown)
		// actuator.POST("/restart", r.actuatorHandler.Restart)
	}

	api := router.Group("/api")
	api.Use(middlewares.AuthMiddleware(r.userRepo, r.cfg, r.log))
	{
		funds := api.Group("/funds")
		{
			funds.Get("", r.fundHandler.GetFunds)
			funds.Post("", r.fundHandler.CreateFund)
			funds.Get("/:fundId", r.fundHandler.GetFund)
			funds.Get("/contributed", r.fundHandler.GetContributedFunds)
		}

		contributions := api.Group("/contributions")
		{
			contributions.Get("", r.contributionHandler.GetContributions)
			contributions.Post("", r.contributionHandler.CreateContribution)
		}
	}
}
