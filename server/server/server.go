package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config *Config
}

func NewServer(c *Config) (*Server, error) {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Server{
		router: NewRouter(&RouterConfig{
			DisableAuth: c.DisableAuth,
		}),
		config: c,
	}, nil
}

func (s *Server) Run() {
	s.router.Run(fmt.Sprintf(":%s", s.config.Port))
}
