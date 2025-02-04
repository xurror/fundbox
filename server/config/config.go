package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// AppConfig contains the configuration for the application
type AppConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Jwt      JwtConfig      `yaml:"jwt"`
	Logger   LoggerConfig   `yaml:"logger"`
	Aws      AwsConfig      `yaml:"aws"`
	Auth0    Auth0Config    `yaml:"auth0"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Sslmode  string `yaml:"sslmode"`
}

type JwtConfig struct {
	SigningKey string        `yaml:"signingKey"`
	Expiration time.Duration `yaml:"expiration"`
}

type LoggerConfig struct {
	Level int32 `yaml:"level"`
}

type AwsConfig struct {
	Region      string         `yaml:"region"`
	Credentials AwsCredentials `yaml:"credentials"`
}

type AwsCredentials struct {
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
}

type Auth0Config struct {
	Domain     string          `yaml:"domain"`
	Audience   string          `yaml:"audience"`
	Management Auth0Management `yaml:"management"`
}

type Auth0Management struct {
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
	Domain       string `yaml:"domain"`
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
