package seeders

import (
	"errors"
	"fmt"
	"user-service/internal/constants"
	"user-service/internal/domain/model"

	"gorm.io/gorm"
)

func seedRoles(tx *gorm.DB) (model.Role, error) {
	roles := []model.Role{
		{
			Code: constants.Admin,
			Name: "Administrator",
		},
		{
			Code: constants.Customer,
			Name: "Customer",
		},
	}

	var AdminRole model.Role

	for _, role := range roles {
		role, err := findOrCreateRole(tx, role)
		if err != nil {
			return model.Role{}, err
		}

		if role.Code == constants.Admin {
			AdminRole = role
		}
	}
	return AdminRole, nil
}

func findOrCreateRole(tx *gorm.DB, roleSeed model.Role) (model.Role, error) {
	var role model.Role

	err := tx.Unscoped().Where("code = ?", roleSeed.Code).First(&role).Error

	switch {
	case err == nil:
		if role.DeletedAt.Valid {
			if err := tx.Unscoped().Model(&role).Update("deleted_at", nil).Error; err != nil {
				return model.Role{}, fmt.Errorf("restore role %q: %w", roleSeed.Code, err)
			}
			role.DeletedAt = gorm.DeletedAt{}
		}
		return role, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		role = model.Role{
			Code: roleSeed.Code,
			Name: roleSeed.Name,
		}
		if err := tx.Create(&role).Error; err != nil {
			return model.Role{}, fmt.Errorf("create role %q: %w", roleSeed.Code, err)
		}
		return role, nil
	default:
		return model.Role{}, fmt.Errorf("find role %q: %w", roleSeed.Code, err)
	}
}
