package utils

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// findProjectRoot searches for the project's root directory by looking for a known marker file.
func ProjectRoot() string {
	// Start from the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Check for a known project marker (adjust based on your project setup)
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir { // If we reach the root directory, stop
			break
		}
		dir = parent
	}

	log.Fatal("project root not found")
	return ""
}

func GetCurrentUserID(c *gin.Context) *uuid.UUID {
	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return nil
	}
	uuidValue := id.(uuid.UUID)
	return &uuidValue
}
