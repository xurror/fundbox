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

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			logger.NewLogger,
			db.NewDatabase,
		),
		fx.Provide(
			repositories.NewFundRepository,
			repositories.NewUserRepository,
		),
		fx.Provide(
			services.NewFundService,
			services.NewUserService,
		),
		fx.Provide(
			handlers.NewFundHandler,
			handlers.NewUserHandler,
		),
		fx.Provide(
			server.NewGinServer, // Initialize Gin server
			routes.NewRouter,    // Register routes
		),
		fx.Invoke(server.StartServer), // Start server
	)

	app.Run()
}
