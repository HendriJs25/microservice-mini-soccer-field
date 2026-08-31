package migration

import (
	"fmt"
	"user-service/internal/config"
	"user-service/internal/database"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func New(cfg config.Database) (*migrate.Migrate, error) {
	m, err := migrate.New("file://migrations", database.BuildPostgresURL(cfg))
	if err != nil {
		return nil, fmt.Errorf("create migration instance: %w", err)
	}

	return m, nil
}
