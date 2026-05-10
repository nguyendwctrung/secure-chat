package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nguyendwctrung/secure-chat/internal/utils"
	ws "github.com/nguyendwctrung/secure-chat/internal/websocket"
)

var Hub = ws.NewHub()

var upgrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c *gin.Context) {
	token := c.Query("token")

	userID, err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}

	client := &ws.Client{
		ID: userID,
		Conn: conn,
		Send: make(chan ws.Message),
		Hub: Hub,
	}

	Hub.Register <- client

	go client.WriteMessage()
	go client.ReadMessage()
}