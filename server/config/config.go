package config

import (
	"community-funds/pkg/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseDSN string // see: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	LogLevel    uint64
	Auth0       Auth0Config
}

type Auth0Config struct {
	Domain   string
	Audience string
}

func envFilePath(filename string) string {
	envPath := filepath.Join(utils.ProjectRoot(), filename)
	return envPath
}

// Helper to read an env var or use default.
func getStringEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

// Helper to read an env var or use default.
func getIntEnv(key string, def int64) int64 {
	parseInt := func(s string) int64 {
		if number, ok := strconv.ParseInt(s, 10, 64); ok == nil {
			return number
		}
		log.Panicf("Error parsing %s as int", s)
		return 0
	}

	if val, ok := os.LookupEnv(key); ok {
		return parseInt(val)
	}
	return def
}

// Helper to read an env var or use default.
func getUintEnv(key string, def uint64) uint64 {
	parseUint := func(s string) uint64 {
		if number, ok := strconv.ParseUint(s, 10, 64); ok == nil {
			return number
		}
		log.Panicf("Error parsing %s as uint", s)
		return 0
	}

	if val, ok := os.LookupEnv(key); ok {
		return parseUint(val)
	}
	return def
}

func NewConfig() *Config {
	envPath := envFilePath(".env")

	// Load environment variables from that file, if present
	if err := godotenv.Load(envPath); err != nil {
		// You could log a warning or return an error, depending on your preference
		log.Printf("Warning: could not load .env from %s: %v", envPath, err)
	}

	// Read each setting or use a default if not set.
	return &Config{
		Port:        getStringEnv("PORT", "8080"),
		LogLevel:    getUintEnv("LOG_LEVEL", 1),
		DatabaseDSN: getStringEnv("DATABASE_DSN", "postgres://psqluser:password@localhost:5432/community_funds?sslmode=disable"),
		Auth0: Auth0Config{
			Domain:   getStringEnv("AUTH0_DOMAIN", ""),
			Audience: getStringEnv("AUTH0_AUDIENCE", ""),
		},
	}
}
