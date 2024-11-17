package main

import (
	"admeliora/chatapp/internal/controller"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	var logger *slog.Logger
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	if os.Getenv("env") == "PROD" {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelWarn,
		}))
	}

	controller := controller.NewController(logger)

	http.HandleFunc("/ws", controller.WebsocketHandler)
	http.ListenAndServe(":8080", nil)
	logger.Info("app ended")
}
