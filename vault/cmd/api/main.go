package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/mrwilliambarenga/t-cube/vault/internal/config"
)

func main() {
	_ = godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()

	slog.Info("starting vault API", "port", cfg.Port)

	select {}
}
