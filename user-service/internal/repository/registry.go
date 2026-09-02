package repository

import (
	"user-service/internal/repository/role"
	"user-service/internal/repository/user"
	"user-service/internal/repository/verificationtoken"

	"gorm.io/gorm"
)

type Registry struct {
	User              user.Repository
	Role              role.Repository
	VerificationToken verificationtoken.Repository
	Transaction       TransactionManager
}

func NewRegistry(db *gorm.DB) *Registry {
	return &Registry{
		User:              user.NewRepository(db),
		Role:              role.NewRepository(db),
		VerificationToken: verificationtoken.NewRepository(db),
		Transaction:       NewTransactionManager(db),
	}
}
