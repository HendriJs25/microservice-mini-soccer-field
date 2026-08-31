package main

import (
	"log/slog"
	"os"
	"time"
	"user-service/config"
	"user-service/database"
	"user-service/handler"
	"user-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("loading config failed", "error", err)
		os.Exit(1)
	}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		slog.Error("loading location failed", "error", err)
		os.Exit(1)
	}
	time.Local = loc

	postgresDB, err := database.NewPostgres(cfg.Database)
	if err != nil {
		slog.Error("connecting to postgres failed", "error", err)
		os.Exit(1)
	}

	defer func() {
		if err := postgresDB.Close(); err != nil {
			slog.Warn("failed to close postgres connection", "error", err)
		}
	}()

	slog.Info("postgres connection established", "host", cfg.Database.Host, "port", cfg.Database.Port, "name", cfg.Database.Name)

	handlerRegistry := handler.NewRegistry()

	router := gin.Default()
	group := router.Group("api/v1")
	routeRegistry := routes.NewRegistry(group, handlerRegistry)
	routeRegistry.Register()

	if err := router.Run(cfg.ServerAddress()); err != nil {
		slog.Error("starting server failed", "error", err)
		os.Exit(1)
	}
}
