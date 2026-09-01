package repository

import (
	"user-service/internal/repository/user"

	"gorm.io/gorm"
)

type Registry struct {
	User user.Repository
}

func NewRegistry(db *gorm.DB) *Registry {
	return &Registry{
		User: user.NewRepository(db),
	}
}
