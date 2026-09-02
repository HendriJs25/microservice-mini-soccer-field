package services

import (
	"user-service/internal/repository"
	"user-service/internal/services/user"
)

type Registry struct {
	User user.Service
}

func NewRegistry(repositories *repository.Registry) *Registry {
	return &Registry{
		User: user.NewService(repositories.User, repositories.Role, repositories.Transaction),
	}
}
