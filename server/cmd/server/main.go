// @title Community Funds API
// @version 1.0
// @description API documentation for the Community Funds system
// @host localhost:8080
// @BasePath /api
package main

import (
	"community-funds/internal/config"
	"community-funds/internal/db"
	"community-funds/internal/handlers"
	"community-funds/internal/repositories"
	"community-funds/internal/routes"
	"community-funds/internal/server"
	"community-funds/internal/services"
	"community-funds/pkg/logger"

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
			services.NewUserService,
		),
		fx.Provide(
			handlers.NewContributionHandler,
			handlers.NewFundHandler,
		),
		fx.Provide(
			server.NewGinServer, // Initialize Gin server
			routes.NewRouter,    // Register routes
		),
		fx.Invoke(server.StartServer), // Start server
	)

	app.Run()
}
