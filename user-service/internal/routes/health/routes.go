package health

import (
	healthhandler "user-service/internal/handler/health"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, handler *healthhandler.Handler) {
	router.GET("/health", handler.Check)
}
