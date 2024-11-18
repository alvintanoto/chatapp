package controller

import (
	"admeliora/chatapp/internal/service"
	"log/slog"
)

type Controller struct {
	logger *slog.Logger

	AuthController AuthController
}

func NewController(logger *slog.Logger, service service.Service) Controller {
	return Controller{
		logger:         logger,
		AuthController: NewAuthController(logger, service),
	}
}

// func (i *implController) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
// 	var upgrader = websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 	}

// 	_, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		i.logger.Error(err.Error())
// 		return
// 	}
// }
