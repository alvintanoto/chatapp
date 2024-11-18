package service

import (
	"admeliora/chatapp/internal/repository"
	"log/slog"
)

type AuthService interface {
	Register(name, email, password string) (err error)
}

type implAuthService struct {
	logger *slog.Logger

	repository repository.Repository
}

func NewAuthService(logger *slog.Logger, repository repository.Repository) AuthService {
	return &implAuthService{logger: logger, repository: repository}
}

func (i *implAuthService) Register(name, email, password string) (err error) {
	// TODO: validate username, email, and password
	err = i.repository.UserRepository.CreateUser(name, email, password)
	if err != nil {
		// TODO: return response error
		return err
	}

	return nil
}
