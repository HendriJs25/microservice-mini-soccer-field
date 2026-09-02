package handler

import (
	healthhandler "user-service/internal/handler/health"
	userhandler "user-service/internal/handler/user"
	"user-service/internal/services"

	customValidator "github.com/go-playground/validator/v10"
)

type Registry struct {
	Health *healthhandler.Handler
	User   *userhandler.Handler
}

func NewRegistry(services *services.Registry, validate *customValidator.Validate) *Registry {
	return &Registry{
		Health: healthhandler.NewHandler(),
		User:   userhandler.NewHandler(services.User, validate),
	}
}
