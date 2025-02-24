package server

import (
	"community-funds/api/routes"
	"community-funds/config"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Server struct {
	app *fiber.App
}

func NewServer(cfg *config.Config, r *routes.Router, log *logrus.Logger) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,https://yourfrontend.com", // Allowed origins
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           int((12 * time.Hour).Seconds()),
	}))

	// app.Use(swagger.New(swagger.Config{
	// 	BasePath: "/swagger", // swagger ui base path
	// 	FilePath: "./docs/swagger.json",
	// }))

	r.SetupRoutes(app) // Register routes

	return &Server{app: app}
}

func StartServer(lc fx.Lifecycle, s *Server, cfg *config.Config, log *logrus.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func(port string) {
				if err := s.app.Listen(":" + port); err != nil {
					log.Fatalf("Server startup error: %v", err)
				}
			}(cfg.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.app.Shutdown()
		},
	})
}
