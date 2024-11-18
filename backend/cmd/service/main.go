package main

import (
	"admeliora/chatapp/internal/controller"
	"admeliora/chatapp/internal/database"
	"admeliora/chatapp/internal/repository"
	"admeliora/chatapp/internal/service"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	router *mux.Router
	db     *pgxpool.Pool

	controller controller.Controller
	service    service.Service
	repository repository.Repository
}

func main() {
	var logger *slog.Logger
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	if os.Getenv("env") == "PROD" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	app := &App{}
	app.db = database.NewDatabase(logger, os.Getenv("database_dsn"))
	app.repository = repository.NewRepository(logger, app.db)
	app.service = service.NewService(logger, app.repository)
	app.controller = controller.NewController(logger, app.service)
	app.router = mux.NewRouter()
	app.setRoutes()

	logger.Info("server starting at port :8080")
	http.ListenAndServe(":8080", app.router)
	logger.Info("app ended")
}
