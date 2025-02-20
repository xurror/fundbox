package server

import (
	"community-funds/internal/config"
	"community-funds/internal/middlewares"
	"community-funds/internal/routes"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Server struct {
	Engine *gin.Engine
}

func NewGinServer(logger *logrus.Logger, cfg *config.Config, r *routes.Router) *Server {
	engine := gin.New()

	// Middleware
	engine.Use(middlewares.GinLogrusMiddleware(logger)) // Attach Logrus middleware
	engine.Use(gin.Recovery())
	engine.Use(middlewares.CorsMiddleware()) // Enable CORS

	r.SetupRoutes(engine) // Register routes

	return &Server{Engine: engine}
}

func StartServer(lc fx.Lifecycle, s *Server, cfg *config.Config, log *logrus.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func(port string) {
				s.Engine.Run(":" + port)
			}(cfg.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// Implement graceful shutdown
			log.Debug("Server shutting down\n")
			return nil
		},
	})
}
