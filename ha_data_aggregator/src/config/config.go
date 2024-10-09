package config

import (
	"fmt"
	"log/slog"

	env "github.com/caarlos0/env/v11"
)

var (
	SERVER_PORT = "SERVER_PORT"
)

type Configuration struct {
	ServerPort    string `env:"SERVER_PORT"`
	PgDatabaseUrl string `env:"POSTGRES_DATABASE_URL"`
	// DatabaseUrl   string `env:"DATABASE_URL"`
	LogLevel  string `env:"LOG_LEVEL"`
	LogFormat string `env:"LOG_FORMAT"`
}

var (
	Config Configuration
)

func Init() error {
	err := env.Parse(&Config)
	if err != nil {
		slog.Error(fmt.Sprintf("Error while parsing environment variables: %s\n", err.Error()))
		return err
	}

	return nil
}
