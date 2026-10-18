package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
	ServerAddr  string `envconfig:"SERVER_ADDR" default:":8080"`
}

func loadConfig() (Config, error) {
	// Existing environment variables take precedence over the optional .env file.
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return Config{}, fmt.Errorf("load .env: %w", err)
	}
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("load config: %w", err)
	}
	if strings.TrimSpace(cfg.DatabaseURL) == "" {
		return Config{}, errors.New("DATABASE_URL must not be empty")
	}
	if strings.TrimSpace(cfg.ServerAddr) == "" {
		return Config{}, errors.New("SERVER_ADDR must not be empty")
	}
	return cfg, nil
}
