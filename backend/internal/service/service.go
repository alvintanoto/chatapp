package service

import (
	"admeliora/chatapp/internal/repository"
	"log/slog"
)

type Service struct {
	AuthService AuthService
}

func NewService(logger *slog.Logger, repository repository.Repository) Service {
	return Service{
		AuthService: NewAuthService(logger, repository),
	}
}
