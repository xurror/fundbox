package server

import (
	"getting-to-go/config"
	appContext "getting-to-go/context"
	"getting-to-go/service"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func NewRouter(
	logger *logrus.Logger,
	config *config.AppConfig,
	authService *service.AuthService,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(RequestLoggerConfig(logger)))
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(CorsConfig()))
	e.Use(middleware.TimeoutWithConfig(TimeoutConfig()))
	e.Use(middleware.BodyLimit("2M"))
	e.Use(appContext.RegisterAppContext())
	//if config.Server.Mode == "debug" {
	//	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	//		logger.WithFields(logrus.Fields{
	//			"request":  string(reqBody),
	//			"response": string(resBody),
	//		}).Debug("REQUEST BODY")
	//	}))
	//}

	api := e.Group("/api", echojwt.WithConfig(JwtConfig(authService)))

	api.GET("/hello", func(ctx echo.Context) error {
		claims := ctx.Get("user")
		logger.Printf("Claims: %v", claims)
		return ctx.JSON(http.StatusOK, claims)
	})

	return e
}
