package config

import (
	"fmt"
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

// findProjectRoot searches for the project's root directory by looking for a known marker file.
func findProjectRoot() (string, error) {
	// Start from the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		// Check for a known project marker (adjust based on your project setup)
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir { // If we reach the root directory, stop
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("project root not found")
}

func envFilePath(filename string) string {
	projectRoot, err := findProjectRoot()
	if err != nil {
		log.Fatalf("Error finding project root: %v", err)
	}

	envPath := filepath.Join(projectRoot, filename)
	return envPath
}

// Helper to read an env var or use default.
func getEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

func parseUint(s string) uint {
	number, _ := strconv.ParseUint(s, 10, 32)
	return uint(number)
}

func NewConfig() *Config {
	// Load environment variables from .env in project root, if it exists.
	envPath := envFilePath(".env")

	// Load environment variables from that file, if present
	if err := godotenv.Load(envPath); err != nil {
		// You could log a warning or return an error, depending on your preference
		log.Printf("Warning: could not load .env from %s: %v", envPath, err)
	}

	// Read each setting or use a default if not set.
	return &Config{
		Port:        getEnv("PORT", "8080"),
		LogLevel:    uint32(parseUint(getEnv("LOG_LEVEL", "1"))),
		DatabaseDSN: getEnv("DATABASE_DSN", "host=localhost user=postgres dbname=postgres sslmode=disable"),
		Auth0: Auth0Config{
			Domain:   getEnv("AUTH0_DOMAIN", ""),
			Audience: getEnv("AUTH0_AUDIENCE", ""),
		},
	}
}
