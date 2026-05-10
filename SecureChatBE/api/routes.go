package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyendwctrung/secure-chat/internal/handlers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	UserRoutes(api)

	router.GET("/ws", handlers.WebSocketHandler)
}