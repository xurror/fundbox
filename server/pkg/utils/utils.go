package utils

import (
	"log"
	"os"
	"path/filepath"
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
