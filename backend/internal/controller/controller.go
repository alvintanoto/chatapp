package controller

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

type Controller interface {
	WebsocketHandler(w http.ResponseWriter, r *http.Request)
}

type implController struct {
	logger *slog.Logger
}

func NewController(logger *slog.Logger) Controller {
	return &implController{
		logger: logger,
	}
}

func (i *implController) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		i.logger.Error(err.Error())
		return
	}
}
