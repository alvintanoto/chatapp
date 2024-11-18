package controller

import (
	"admeliora/chatapp/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
)

type AuthController interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type implAuthController struct {
	logger  *slog.Logger
	service service.Service
}

func NewAuthController(logger *slog.Logger, service service.Service) AuthController {
	return &implAuthController{logger: logger, service: service}
}

func (i *implAuthController) Register(w http.ResponseWriter, r *http.Request) {
	type RegisterDTO struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	payload := &RegisterDTO{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		i.logger.Error(err.Error())

		sendResponse(w, r, http.StatusBadRequest, Response{
			Code:    "100000",
			Message: "failed to decode payload",
		})
		return
	}

	err = i.service.AuthService.Register(payload.Name, payload.Email, payload.Password)
	if err != nil {
		// TODO: return response error
		sendResponse(w, r, http.StatusInternalServerError, Response{
			Code:    "200000",
			Message: "failed to register user",
		})
		return
	}

	// TODO: return success
	sendResponse(w, r, http.StatusOK, Response{
		Code:    "000000",
		Message: "success",
	})
}
