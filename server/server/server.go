package server

import (
	"context"
	"fmt"
	"getting-to-go/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

func NewServer(lc fx.Lifecycle, e *echo.Echo, c *config.AppConfig) *http.Server {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go e.Start(fmt.Sprintf(":%s", c.Server.Port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
	return e.Server
}
