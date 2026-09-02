package verificationtoken

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
	Create(context.Context, *model.VerificationToken) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, token *model.VerificationToken) error {
	if err := r.db.WithContext(ctx).Create(token).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errConstant.ErrAlreadyExists
		}
		return fmt.Errorf("create verification token: %w", err)
	}
	return nil
}
