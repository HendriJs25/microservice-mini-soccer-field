package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"user-service/internal/config"
	"user-service/internal/database"
	"user-service/internal/database/seeders"

	"github.com/spf13/cobra"
)

const seedTimeout = 30 * time.Second

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run seeders",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runSeeder()
	},
}

func runSeeder() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("loading config failed: %w", err)
	}

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

	if err := cfg.Seed.Validate(); err != nil {
		return fmt.Errorf("validate seed configuration: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), seedTimeout)
	defer cancel()

	if err := seeders.Run(ctx, postgresDB.DB, cfg.Seed); err != nil {
		return err
	}

	slog.Info("seed initialized successfully")

	return nil
}
