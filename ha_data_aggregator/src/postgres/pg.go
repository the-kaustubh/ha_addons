package pg

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const (
	TEMPERATURE_TABLE_NAME = "public.pitm"
	// INSERT INTO public.pitm (machine, temperature) VALUES('rpizw', 49.3);
)

func Init(ctx context.Context, connectionString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx,
		connectionString,
	)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		log.Println("Error while pinging the psql db: " + err.Error())
		return nil, err
	}
	return conn, nil
}

func WriteTemperature(ctx context.Context, conn *pgx.Conn, machineName string, temperature float64) error {
	query := fmt.Sprintf("INSERT INTO %s (machine, temperature) VALUES(@machine, @temperature);", TEMPERATURE_TABLE_NAME)
	_, err := conn.Exec(ctx, query,
		pgx.NamedArgs{
			"machine":     machineName,
			"temperature": temperature,
		},
	)
	if err != nil {
		log.Println("Error while writing temperature to psql db: " + err.Error())
		return err
	}
	return nil
}
