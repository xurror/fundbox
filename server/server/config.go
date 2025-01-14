package server

import (
	"getting-to-go/service"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func CorsConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
}

func TimeoutConfig() middleware.TimeoutConfig {
	return middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "request timeout",
		Timeout:      30 * time.Second,
	}
}

func RequestLoggerConfig(log *logrus.Logger) middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			status := values.Status
			var logLevel logrus.Level

			if status >= 200 && status < 300 {
				logLevel = logrus.InfoLevel
			} else if status >= 300 && status < 400 {
				logLevel = logrus.WarnLevel
			} else if status >= 400 && status < 500 {
				logLevel = logrus.DebugLevel
			} else {
				logLevel = logrus.ErrorLevel
			}

			log.WithFields(logrus.Fields{
				"method": values.Method,
				"URI":    values.URI,
				"status": status,
			}).Log(logLevel, "request")
			return nil
		},
	}
}

func JwtConfig(authService *service.AuthService) echojwt.Config {
	return echojwt.Config{
		ContextKey:     "user",
		ParseTokenFunc: authService.Authorize,
	}
}
