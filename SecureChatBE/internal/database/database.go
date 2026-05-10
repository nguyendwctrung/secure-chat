package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nguyendwctrung/secure-chat/internal/config"
	"github.com/nguyendwctrung/secure-chat/internal/models"
)

var DB *gorm.DB

func ConnectDatabase(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	DB = database

	log.Println("Database connected successfully")

	err = DB.AutoMigrate(&models.User{})
	
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}