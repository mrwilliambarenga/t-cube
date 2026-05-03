package config

import (
	"os"
)

type Config struct {
	Port string
	Env  string
}

func Load() Config {
	cfg := Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("ENV", "dev"),
	}

	return cfg
}

func getEnv(key string, fallback string) string {
	val, exists := os.LookupEnv(key)

	if !exists || val == "" {
		return fallback
	}

	return val
}
