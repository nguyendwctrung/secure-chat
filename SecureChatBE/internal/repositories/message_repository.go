package repositories

import (
	"github.com/nguyendwctrung/secure-chat/internal/database"
	"github.com/nguyendwctrung/secure-chat/internal/models"
)

func SaveMessage(message *models.Message) error {
	return database.DB.Create(message).Error
}