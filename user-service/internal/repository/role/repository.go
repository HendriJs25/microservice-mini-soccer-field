package role

import (
	"context"
	"errors"
	"fmt"
	errConstant "user-service/internal/constants/error"
	"user-service/internal/domain/model"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindByCode(context.Context, string) (*model.Role, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindByCode(ctx context.Context, code string) (*model.Role, error) {
	var role model.Role
	if err := r.db.WithContext(ctx).Where("code = ?", code).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("find role by code : %w", errConstant.ErrNotFound)
		}
		return nil, fmt.Errorf("find role by code : %w", err)
	}
	return &role, nil
}
