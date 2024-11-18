package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database *pgxpool.Pool

func NewDatabase(logger *slog.Logger, dbUrl string) Database {
	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		logger.Info(err.Error())
	}

	return dbpool
}
