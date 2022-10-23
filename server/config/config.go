package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// Database configuration
	DbUser     string `envconfig:"DB_USER" default:"root"`
	DbPassword string `envconfig:"DB_PASSWORD" default:""`
	DbName     string `envconfig:"DB_NAME" default:"items_db"`
	DbHost     string `envconfig:"DB_HOST" default:"localhost"`

	// Server configuration
	ServerPort string `envconfig:"SERVER_PORT" default:"8080"`
}

// GetConfig returns the configuration of the application
func GetConfig() *Config {
	var cfg Config
	envconfig.Process("", &cfg)
	return &cfg
}
