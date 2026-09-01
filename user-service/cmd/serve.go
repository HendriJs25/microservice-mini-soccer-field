package cmd

import (
	"fmt"
	"log/slog"
	"time"
	"user-service/internal/config"
	"user-service/internal/database"
	"user-service/internal/handler"
	"user-service/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server",

	RunE: func(cmd *cobra.Command, args []string) error {
		return runServer()
	},
}

func runServer() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("loading config failed: %w", err)
	}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return fmt.Errorf("loading location failed: %w", err)
	}
	time.Local = loc

	postgresDB, err := database.NewPostgres(cfg.Database)
	if err != nil {
		return fmt.Errorf("connecting to postgres failed: %w", err)
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
		return fmt.Errorf("starting server failed: %w", err)
	}
	return nil
}
