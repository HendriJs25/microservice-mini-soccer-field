package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"user-service/internal/config"
	"user-service/internal/database/migration"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Manage database migrations",
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Run pending database migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMigrateUp()
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback last database migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMigrateDown()
	},
}

var migrateForceCmd = &cobra.Command{
	Use:   "force [version]",
	Short: "force database migration version [-1 for no migrations]",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		version, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid migration version: %w", err)
		}
		return runMigrateForce(version)
	},
}

var migrateVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current database version",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMigrateVersion()
	},
}

func runMigrateUp() error {
	m, err := createMigrateInstance()
	if err != nil {
		return err
	}
	defer closeInstance(m)

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			slog.Info("Nothing to migrate")
			return nil
		}
		return fmt.Errorf("migrate up: %w", err)
	}
	slog.Info("Migration completed successfully")
	return nil
}

func runMigrateDown() error {
	m, err := createMigrateInstance()
	if err != nil {
		return err
	}
	defer closeInstance(m)

	if err := m.Steps(-1); err != nil {
		return fmt.Errorf("migrate down: %w", err)
	}

	slog.Info("Migration rollback completed")
	return nil
}

func runMigrateForce(version int) error {
	m, err := createMigrateInstance()
	if err != nil {
		return err
	}
	defer closeInstance(m)

	if err := m.Force(version); err != nil {
		return fmt.Errorf("force migration version: %w", err)
	}

	slog.Info("Migration version forced successfully", "version", version)
	return nil
}

func runMigrateVersion() error {
	m, err := createMigrateInstance()
	if err != nil {
		return err
	}
	defer closeInstance(m)

	version, dirty, err := m.Version()
	if err != nil {
		if errors.Is(err, migrate.ErrNilVersion) {
			slog.Info("No migration has been applied")
			return nil
		}
		return fmt.Errorf("get migration version: %w", err)
	}

	slog.Info("Current migration version",
		"version", version,
		"dirty", dirty)

	return nil
}

func createMigrateInstance() (*migrate.Migrate, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}
	m, err := migration.New(cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("initialize migration: %w", err)
	}
	return m, nil
}

func closeInstance(m *migrate.Migrate) {
	sourceErr, databaseErr := m.Close()
	if sourceErr != nil {
		slog.Warn("Failed to close migration source", "error", sourceErr)
	}
	if databaseErr != nil {
		slog.Warn("Failed to close migration database", "error", databaseErr)
	}
}
