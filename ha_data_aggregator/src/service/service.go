package service

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/the-kaustubh/ha_data_aggregator/config"
	hacontext "github.com/the-kaustubh/ha_data_aggregator/context"
	pg "github.com/the-kaustubh/ha_data_aggregator/postgres"
)

var (
	configuration config.Configuration
	conn          *pgx.Conn
)

func Init(config config.Configuration) error {
	configuration = config
	var err error
	slog.Info("Attempting to connect to pgsql database")
	conn, err = pg.Init(context.Background(), config.PgDatabaseUrl)
	if err != nil {
		slog.Error("Error while connecting to postgres database", "error", err.Error())
		return err
	}
	slog.Info("Attempting to connect to pgsql database: done")
	return nil
}

func WriteTemperature(ctx hacontext.Context, machineName string, temperature float64) error {
	slog.Debug("Writing temperature", "machineName", machineName, "teperature", temperature)
	return pg.WriteTemperature(ctx, conn, machineName, temperature)
}
