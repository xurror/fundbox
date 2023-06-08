package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// AppConfig contains the configuration for the application
type AppConfig struct {
	Server struct {
		Port string `yaml:"port"`
	}
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

// LoadConfig loads the application configuration from the given YAML file
func LoadConfig(filename string) (*AppConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	// expand environment variables
	data = []byte(os.ExpandEnv(string(data)))
	conf := &AppConfig{}
	if err := yaml.Unmarshal(data, conf); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return conf, nil
}
