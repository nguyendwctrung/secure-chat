package main

import (
	"github.com/gin-gonic/gin"

	"github.com/nguyendwctrung/secure-chat/internal/config"
	"github.com/nguyendwctrung/secure-chat/internal/database"
	"github.com/nguyendwctrung/secure-chat/api"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDatabase(cfg)

	router := gin.Default()

	api.RegisterRoutes(router)

	router.Run(":" + cfg.Port)
}