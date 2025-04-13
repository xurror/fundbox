// @title Community Funds API
// @version 1.0
// @description API documentation for the Community Funds system
// @host localhost:8080
// @BasePath /api
package main

import (
	"community-funds/api/handlers"
	"community-funds/api/routes"
	"community-funds/api/server"
	"community-funds/config"
	"community-funds/logger"
	"community-funds/pkg/db"
	"community-funds/pkg/repositories"
	"community-funds/pkg/services"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	app := fx.New(
		fx.WithLogger(func(logrus *logrus.Logger) fxevent.Logger {
			return &logger.Logger{Logger: logrus}
		}),
		fx.Provide(
			config.NewConfig,
			logger.NewLogger,
			db.NewDatabase,
		),
		fx.Provide(
			repositories.NewContributionRepository,
			repositories.NewFundRepository,
			repositories.NewUserRepository,
		),
		fx.Provide(
			services.NewContributionService,
			services.NewFundService,
			services.NewStripeService,
			services.NewUserService,
		),
		fx.Provide(
			handlers.NewActuatorHandler,
			handlers.NewContributionHandler,
			handlers.NewFundHandler,
			handlers.NewStripeHandler,
			handlers.NewUserHandler,
		),
		fx.Provide(
			routes.NewRouter, // Register routes
			server.NewServer, // Initialize App server
		),
		fx.Invoke(server.StartServer), // Start server
	)

	app.Run()
}
