package server

import (
	"community-funds/api/routes"
	"community-funds/config"
	"community-funds/pkg/utils"
	"context"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
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
	// fiberlog.SetLogger(log)
	fiberlog.SetLevel(fiberlog.Level(cfg.LogLevel))
	app := fiber.New(fiber.Config{
		AppName:         "Community Funds",
		ReadBufferSize:  8192,
		WriteBufferSize: 8192,
	})

	// app.Use(func(c *fiber.Ctx) error {
	// 	fmt.Printf("Request URL: %s\n", c.OriginalURL())
	// 	headerSize := 0
	// 	for key, values := range c.GetReqHeaders() {
	// 		for _, value := range values {
	// 			fmt.Printf("Request Header Key: %s\n", key)
	// 			headerSize += len(key) + len(value)
	// 		}
	// 	}

	// 	fmt.Printf("Request header size: %d bytes\n", headerSize)
	// 	return c.Next()
	// })

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cfg.Cors.ToServerCors()))

	app.Use(swagger.New(swagger.Config{
		BasePath: "/swagger", // swagger ui base path
		FilePath: utils.GetFilePath("docs/swagger.json"),
	}))

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
