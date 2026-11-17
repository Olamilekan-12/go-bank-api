package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	Env  string
}

func Load() (Config, error) {
	cfg := Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("APP_ENV", "development"),
	}

	if cfg.Env != "development" && cfg.Env != "production" {
		return Config{}, fmt.Errorf("invalid APP_ENV %q: must be development or production", cfg.Env)
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
