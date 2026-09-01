package user

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
	Create(context.Context, *model.User) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, user *model.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("create user: %w", errConstant.ErrAlreadyExists)
		}
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}
