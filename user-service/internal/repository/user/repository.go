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
	ExistByEmail(context.Context, string) (bool, error)
	MarkVerified(context.Context, int64) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, user *model.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errConstant.ErrAlreadyExists
		}
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (r *repository) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, fmt.Errorf("is user exist: %w", err)
	}
	return count > 0, nil
}

func (r *repository) MarkVerified(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("is_verified", true)

	if result.Error != nil {
		return fmt.Errorf("mark verified user: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errConstant.ErrNotFound
	}

	return nil
}
