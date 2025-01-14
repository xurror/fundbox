package main

import (
	appConfig "getting-to-go/config"
	"getting-to-go/controller"
	"getting-to-go/logging"
	"getting-to-go/model"
	"getting-to-go/server"
	"getting-to-go/service"
	"net/http"

	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	app := fx.New(
		fx.WithLogger(func(logger *log.Logger) fxevent.Logger {
			return &logging.AppLogger{Logger: logger}
		}),
		fx.Provide(
			logging.NewLogger,
			appConfig.NewAppConfig,
			model.NewDynamoDBClient,
		),
		fx.Provide(
			service.NewUserServiceDynamoDbImpl,
			service.NewAuthService,
		),
		fx.Provide(
			controller.NewUserController,
		),
		fx.Provide(
			server.NewServer,
			server.NewRouter,
		),
		fx.Invoke(func(*http.Server) {}),
	)

	app.Run()
}
