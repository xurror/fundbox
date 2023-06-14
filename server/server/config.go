package server

import (
	"getting-to-go/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
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

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "*"
		//},
	}
}

func SecureConfig() secure.Config {
	return secure.Config{
		//SSLRedirect:           false,
		//IsDevelopment:         false,
		//STSSeconds:            315360000,
		//STSIncludeSubdomains:  false,
		//FrameDeny:             false,
		//ContentTypeNosniff:    false,
		//BrowserXssFilter:      false,
		//ContentSecurityPolicy: "",
		//IENoOpen:              false,
		//SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}
}
