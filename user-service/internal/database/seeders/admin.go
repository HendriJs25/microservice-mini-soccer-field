package seeders

import (
	"errors"
	"fmt"
	"user-service/internal/config"
	"user-service/internal/domain/model"
	"uuid"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seedAdmin(tx *gorm.DB, adminRole model.Role, cfg config.Seed) error {
	var admin model.User

	err := tx.Unscoped().Where("email = ?", cfg.AdminEmail).First(&admin).Error

	switch {
	case err == nil:
		updates := map[string]any{}

		if admin.DeletedAt.Valid {
			updates["deleted_at"] = nil
		}

		if !admin.IsVerified {
			updates["is_verified"] = true
		}

		if len(updates) > 0 {
			if err := tx.Unscoped().Model(&admin).Updates(updates).Error; err != nil {
				return fmt.Errorf("restore admin user: %w", err)
			}
		}

	case errors.Is(err, gorm.ErrRecordNotFound):
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("generate password hash: %w", err)
		}

		admin = model.User{
			UUID:         uuid.New(),
			Name:         "Admin",
			Username:     "Admin",
			PasswordHash: string(passwordHash),
			Email:        cfg.AdminEmail,
			IsVerified:   true,
			RoleID:       adminRole.ID,
		}

		if err := tx.Create(&admin).Error; err != nil {
			return fmt.Errorf("create admin user: %w", err)
		}

	default:
		return fmt.Errorf("find admin user: %w", err)
	}
	return nil
}
