package user

import (
	userhandler "user-service/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, handler *userhandler.Handler) {
	router.POST("/register", handler.Register)
}
