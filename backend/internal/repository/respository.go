package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	logger *slog.Logger

	UserRepository UserRepository
}

func NewRepository(logger *slog.Logger, db *pgxpool.Pool) Repository {
	return Repository{
		logger:         logger,
		UserRepository: NewUserRepository(logger, db),
	}
}
