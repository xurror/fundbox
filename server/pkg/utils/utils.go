package utils

import (
	"log"
	"os"
	"path/filepath"
)

func logDirContents(dirPath string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Println("Error reading directory:", err)
	}

	log.Println("Directory contents:")
	for _, entry := range entries {
		log.Println(entry.Name()) // Prints file/directory names
	}
}

func GetFilePath(filename string) string {
	projectRoot := ProjectRoot()
	// logDirContents(projectRoot)

	envPath := filepath.Join(projectRoot, filename)
	return envPath
}

// findProjectRoot searches for the project's root directory by looking for a known marker file.
func ProjectRoot() string {
	projectRoot := os.Getenv("PROJECT_ROOT_DIR")
	if projectRoot != "" {
		return projectRoot
	}

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
	return ""
}
