package config

import (
	"community-funds/pkg/utils"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type Config struct {
	StripeKey   string
	Port        string
	DatabaseDSN string
	LogLevel    uint64
	Auth0       *Auth0Config
	Cors        *CorsConfig
}

type CorsConfig struct {
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	ExposeHeaders    string
	AllowCredentials bool
	MaxAge           int
}

type Auth0Config struct {
	Domain   string
	Audience string
}

func NewConfig() *Config {
	envPath := utils.GetFilePath(".env")

	// Load environment variables from that file, if present
	if err := godotenv.Load(envPath); err != nil {
		// You could log a warning or return an error, depending on your preference
		log.Printf("Warning: could not load .env from %s: %v", envPath, err)
	}

	// Read each setting or use a default if not set.
	return &Config{
		StripeKey:   getStringEnv("STRIPE_KEY", ""),
		Port:        getStringEnv("PORT", "8080"),
		LogLevel:    getUintEnv("LOG_LEVEL", 1),
		DatabaseDSN: getStringEnv("DATABASE_DSN", "postgres://psqluser:password@localhost:5432/community_funds?sslmode=disable"),
		Auth0: &Auth0Config{
			Domain:   getStringEnv("AUTH0_DOMAIN", ""),
			Audience: getStringEnv("AUTH0_AUDIENCE", ""),
		},
		Cors: &CorsConfig{
			AllowOrigins:     getStringEnv("CORS_ALLOW_ORIGINS", "http://localhost:3000"),
			AllowMethods:     getStringEnv("CORS_ALLOW_METHODS", "GET,POST,PUT,DELETE,OPTIONS"),
			AllowHeaders:     getStringEnv("CORS_ALLOW_HEADERS", "Origin,Content-Type,Authorization"),
			ExposeHeaders:    getStringEnv("CORS_EXPOSE_HEADERS", "Content-Length"),
			AllowCredentials: getBoolEnv("CORS_ALLOW_CREDENTIALS", true),
			MaxAge:           getIntEnv("CORS_ALLOW_ORIGINS", int((12 * time.Hour).Seconds())),
		},
	}
}

func (c *CorsConfig) ToServerCors() cors.Config {
	return cors.Config{
		AllowOrigins:     c.AllowOrigins,
		AllowMethods:     c.AllowMethods,
		AllowHeaders:     c.AllowHeaders,
		ExposeHeaders:    c.ExposeHeaders,
		AllowCredentials: c.AllowCredentials,
		MaxAge:           c.MaxAge,
	}
}

func getStringEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

func getBoolEnv(key string, def bool) bool {
	parseBool := func(s string) bool {
		if val, ok := strconv.ParseBool(s); ok == nil {
			return val
		} else {
			return def
		}
	}

	if val, ok := os.LookupEnv(key); ok {
		return parseBool(val)
	}
	return def
}

func getIntEnv(key string, def int) int {
	parseInt := func(s string) int {
		if number, ok := strconv.ParseInt(s, 10, 64); ok == nil {
			return int(number)
		}
		return def
	}

	if val, ok := os.LookupEnv(key); ok {
		return parseInt(val)
	}
	return def
}

func getUintEnv(key string, def uint64) uint64 {
	parseUint := func(s string) uint64 {
		if number, ok := strconv.ParseUint(s, 10, 64); ok == nil {
			return number
		}
		return def
	}

	if val, ok := os.LookupEnv(key); ok {
		return parseUint(val)
	}
	return def
}
