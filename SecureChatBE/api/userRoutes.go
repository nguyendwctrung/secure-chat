package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyendwctrung/secure-chat/internal/handlers"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	{
		users.GET("/health", handlers.HealthCheck)
	}
}