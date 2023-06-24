package server

import (
	"getting-to-go/config"
	"github.com/go-chi/cors"
)

type Config struct {
	Port        string
	Debug       bool
	DisableAuth bool
}

func NewConfig(config *config.AppConfig) *Config {
	return &Config{
		Port:        config.Server.Port,
		Debug:       config.Server.Debug,
		DisableAuth: config.Server.DisableAuth,
	}
}

type RouterConfig struct {
	DisableAuth bool
}

func CorsOptions() cors.Options {
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	return cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
}
