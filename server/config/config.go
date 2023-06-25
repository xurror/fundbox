package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

// AppConfig contains the configuration for the application
type AppConfig struct {
	Server struct {
		Port        string `yaml:"port"`
		Debug       bool   `yaml:"debug"`
		DisableAuth bool   `yaml:"disableAuth"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`

	Jwt struct {
		SigningKey string        `yaml:"signingKey"`
		Expiration time.Duration `yaml:"expiration"`
	} `yaml:"jwt"`
}

func NewAppConfig() *AppConfig {
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Panic(err.Error())
	}

	var config AppConfig
	data = []byte(os.ExpandEnv(string(data)))
	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Panic(err.Error())
	}
	return &config
}
