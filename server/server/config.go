package server

import (
	"getting-to-go/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"time"
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
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "*"
		//},
		MaxAge: 12 * time.Hour,
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
