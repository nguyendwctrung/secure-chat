package main

import (
	"github.com/gin-gonic/gin"

	"github.com/nguyendwctrung/secure-chat/api"
	"github.com/nguyendwctrung/secure-chat/internal/config"
	"github.com/nguyendwctrung/secure-chat/internal/database"
	"github.com/nguyendwctrung/secure-chat/internal/handlers"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDatabase(cfg)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	api.SetupRoutes(router)
	go handlers.Hub.Run()

	router.Run(":" + cfg.Port)
}