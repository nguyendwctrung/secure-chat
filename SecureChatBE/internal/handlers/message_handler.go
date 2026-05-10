package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyendwctrung/secure-chat/internal/database"
	"github.com/nguyendwctrung/secure-chat/internal/models"
)

func GetMessage(c *gin.Context) {
	senderID := c.Query("sender_id")
	receiverID := c.Query("receiver_id")

	var messages []models.Message

	database.DB.Where(
		"(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		senderID,
		receiverID,
		receiverID,
		senderID,
	).Order("created_at asc").Find(&messages)

	c.JSON(http.StatusOK, messages)
}