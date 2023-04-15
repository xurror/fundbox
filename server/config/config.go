package config

import (
	"fmt"
	"io/ioutil"

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
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &cfg, nil
}
