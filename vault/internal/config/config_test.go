package config

import "testing"

func TestGetEnv(t *testing.T) {
	t.Setenv("PORT", "9999")

	val := getEnv("PORT", "fallback")

	if val != "9999" {
		t.Errorf("Expected 9999, got %s", val)
	}
}

func TestGetEnvFallback(t *testing.T) {
	val := getEnv("PORT", "fallback")

	if val != "fallback" {
		t.Errorf("Expected fallback, got %s", val)
	}
}
