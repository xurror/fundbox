package server

import (
	"community-funds/internal/config"
	"community-funds/internal/middlewares"
	"community-funds/internal/routes"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Server struct {
	Engine *gin.Engine
}

func NewGinServer(logger *logrus.Logger, cfg *config.Config, r *routes.Router) *Server {
	// gin.SetMode(gin.ReleaseMode) // Set Gin to release mode for performance
	gin.SetMode(gin.DebugMode) // Set Gin to release mode for performance
	engine := gin.Default()    // Includes logging and recovery middleware

	r.SetupRoutes(engine) // Register routes

	// Middleware

	// Attach Logrus middleware
	engine.Use(middlewares.GinLogrusMiddleware(logger))
	engine.Use(gin.Recovery())

	return &Server{Engine: engine}
}

func StartServer(lc fx.Lifecycle, s *Server, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func(port string) {
				if port == "" {
					port = "8080"
				}
				fmt.Printf("Server running on port %s\n", port)

				s.Engine.Run(":" + port)
			}(cfg.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// Implement graceful shutdown
			fmt.Printf("Server shutting down\n")
			return nil
		},
	})
}
