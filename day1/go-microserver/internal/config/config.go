package config

import (
	"os"
)

type Config struct {
	Port     string
	LogLevel string
}

func Load() Config {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // default
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "INFO"
	}

	return Config{
		Port:     port,
		LogLevel: logLevel,
	}
}
