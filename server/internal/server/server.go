package server

import (
	"community-funds/internal/config"
	"community-funds/internal/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func NewGinServer(cfg *config.Config) *Server {
	gin.SetMode(gin.ReleaseMode) // Set Gin to release mode for performance
	router := gin.Default()      // Includes logging and recovery middleware

	return &Server{Router: router}
}

func StartServer(s *Server, cfg *config.Config, r *routes.Router) {
	r.SetupRoutes(s.Router) // Register routes
	port := cfg.Port
	fmt.Printf("Server running on port %s\n", port)
	s.Router.Run(":" + port)
}
