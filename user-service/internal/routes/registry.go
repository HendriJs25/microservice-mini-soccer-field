package routes

import (
	"user-service/internal/handler"
	healthroutes "user-service/internal/routes/health"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	router   *gin.RouterGroup
	handlers *handler.Registry
}

func NewRegistry(router *gin.RouterGroup, handler *handler.Registry) *Registry {
	return &Registry{
		router:   router,
		handlers: handler,
	}
}

func (r *Registry) Register() {
	healthroutes.Register(r.router, r.handlers.Health)
}
