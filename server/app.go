package main

import (
	"getting-to-go/config"
	"getting-to-go/controller"
	"getting-to-go/logging"
	"getting-to-go/model"
	"getting-to-go/server"
	"getting-to-go/service"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"net/http"
)

func main() {
	app := fx.New(
		fx.WithLogger(func(logger *log.Logger) fxevent.Logger {
			return &logging.AppLogger{Logger: logger}
		}),
		fx.Provide(
			model.NewDB,
			logging.NewLogger,
			config.NewAppConfig,
		),
		fx.Provide(
			service.NewUserService,
			service.NewAuthService,
			service.NewContributionService,
			service.NewFundService,
		),
		fx.Provide(
			controller.NewAuthController,
		),
		fx.Provide(
			server.NewGraphQlHandler,
			server.NewServer,
			server.NewRouter,
		),
		fx.Invoke(func(*http.Server) {}),
	)

	app.Run()
}
