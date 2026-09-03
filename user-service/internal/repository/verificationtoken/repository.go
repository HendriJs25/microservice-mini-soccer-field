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
	FindByHashToken(context.Context, string) (*model.VerificationToken, error)
	Delete(context.Context, string) error
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

func (r *repository) FindByHashToken(ctx context.Context, hashToken string) (*model.VerificationToken, error) {
	var token model.VerificationToken

	if err := r.db.WithContext(ctx).Where("hash_token = ?", hashToken).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConstant.ErrNotFound
		}
		return nil, fmt.Errorf("find verification token by hash token: %w", err)
	}

	return &token, nil
}

func (r *repository) Delete(ctx context.Context, hashToken string) error {
	result := r.db.WithContext(ctx).Where("hash_token = ?", hashToken).Delete(&model.VerificationToken{})

	if result.Error != nil {
		return fmt.Errorf("delete verification token: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errConstant.ErrNotFound
	}
	return nil
}
