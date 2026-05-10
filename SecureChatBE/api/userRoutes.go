package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyendwctrung/secure-chat/internal/handlers"
	"github.com/nguyendwctrung/secure-chat/internal/middleware"
)

func UserRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")

	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.GET("/profile", middleware.AuthMiddleware(), handlers.Profile,)
		auth.GET("/messages", middleware.AuthMiddleware(), handlers.GetMessage,)
	}
}