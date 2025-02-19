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
	LogLevel    uint32
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
func getEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

func parseUint(s string, defaultVal int) uint {
	number, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return uint(defaultVal)
	}
	return uint(number)
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
		Port:        getEnv("PORT", "8080"),
		LogLevel:    uint32(parseUint(getEnv("LOG_LEVEL", "1"), 1)),
		DatabaseDSN: getEnv("DATABASE_DSN", "host=localhost user=postgres dbname=postgres sslmode=disable"),
		Auth0: Auth0Config{
			Domain:   getEnv("AUTH0_DOMAIN", ""),
			Audience: getEnv("AUTH0_AUDIENCE", ""),
		},
	}
}
