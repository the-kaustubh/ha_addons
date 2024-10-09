package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/the-kaustubh/ha_data_aggregator/config"
	"github.com/the-kaustubh/ha_data_aggregator/server"
	"github.com/the-kaustubh/ha_data_aggregator/service"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/ha_data_aggregator.log",
		MaxSize:    4, // 4MB
		MaxAge:     28,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   true,
	}

	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, lumberjackLogger), &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	err := config.Init()
	Must(err)

	logLevel := &slog.LevelVar{}
	err = logLevel.UnmarshalText([]byte(config.Config.LogLevel))
	Must(err)
	slog.SetLogLoggerLevel(logLevel.Level())
	if config.Config.LogFormat == "text" {
		textLogger := slog.NewTextHandler(io.MultiWriter(os.Stdout, lumberjackLogger), &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
		slog.SetDefault(slog.New(textLogger))
	}

	slog.Info("Initiating service")
	Must(service.Init(config.Config))
	slog.Info("Initiated service")
	Must(server.Init(config.Config))
}

func Must(err error) {
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
