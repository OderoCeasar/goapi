package config

import (
	"fmt"
	"os"
)

// config holds application configuration
type Config struct {
	Port string
	Env	string
	AppName	string
	Version string

}

// load reads configuration from env 
func Load() (*Config, error) {
	cfg := &Config{
		Port: 		getEnv("PORT", "8080"),
		Env: 		getEnv("APP_ENV", "production"),
		AppName: 	getEnv("APP_NAME", "goapi"),
		Version: 	getEnv("APP_VERSION", "1.0.0"),
	}

	// validate required fields
	if cfg.Port == "" {
		return nil, fmt.Errorf("PORT environment variable is required")

	}

	return cfg, nil
}

// getEnv reads an env var and returns a fallback if not set
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}