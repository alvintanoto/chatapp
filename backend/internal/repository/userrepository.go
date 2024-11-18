package repository

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(name, email, password string) (err error)
}

type implUserRepository struct {
	logger *slog.Logger
	db     *pgxpool.Pool
}

func NewUserRepository(logger *slog.Logger, db *pgxpool.Pool) UserRepository {
	return &implUserRepository{
		logger: logger,
		db:     db,
	}
}

func (i *implUserRepository) CreateUser(name, email, password string) (err error) {
	userID := uuid.NewString()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		i.logger.Error("error hashing user password", "description", err.Error())
		return err
	}

	query := `INSERT INTO public.user(id, name, email, password) VALUES ($1, $2, $3, $4)`
	args := []interface{}{userID, name, email, string(hashedPassword)}

	_, err = i.db.Exec(context.Background(), query, args...)
	if err != nil {
		i.logger.Error("error creating user", "description", err.Error())
		return err
	}

	return nil
}
