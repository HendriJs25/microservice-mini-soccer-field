package seeders

import (
	"context"
	"fmt"
	"user-service/internal/config"

	"gorm.io/gorm"
)

func Run(ctx context.Context, db *gorm.DB, cfg config.Seed) error {
	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		adminRole, err := seedRoles(tx)
		if err != nil {
			return fmt.Errorf("seed roles: %w", err)
		}

		if err := seedAdmin(tx, adminRole, cfg); err != nil {
			return fmt.Errorf("seed admin: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("run seeders: %w", err)
	}

	return nil
}
