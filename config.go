package memogram

import (
	"context"
	"log/slog"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	ServerAddr string `env:"SERVER_ADDR,required"`
	BotToken   string `env:"BOT_TOKEN,required"`
}

func getConfigFromEnv() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Log(context.Background(), slog.LevelError, "failed to load env", "msg", err.Error())
	}

	config := Config{}
	if err := env.Parse(&config); err != nil {
		return nil, errors.Wrap(err, "invalid configuration")
	}
	return &config, nil
}
