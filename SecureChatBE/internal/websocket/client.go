package websocket

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/nguyendwctrung/secure-chat/internal/models"
	"github.com/nguyendwctrung/secure-chat/internal/repositories"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan Message
	Hub *Hub
}

func (c *Client) ReadMessage() {
	defer func ()  {
		c.Hub.Unregister <- c
		c.Conn.Close()
	} ()

	for {
		var msg Message

		err := c.Conn.ReadJSON(&msg)

		if err != nil {
			log.Println("Read error:", err)
			break
		}

		msg.SenderID = c.ID
		senderUUID, _ := uuid.Parse(msg.SenderID)
		receverUUID, _ := uuid.Parse(msg.ReceiverID)

		dbMessage := models.Message{
			ID: uuid.New(),
			SenderID: senderUUID,
			ReceiverID: receverUUID,
			Content: msg.Content,
		}

		repositories.SaveMessage(&dbMessage)

		c.Hub.Broadcast <- msg
	}
}

func (c *Client) WriteMessage() {
	defer c.Conn.Close()

	for {
		msg, ok := <- c.Send

		if !ok {
			return
		}

		err := c.Conn.WriteJSON(msg)

		if err != nil {
			log.Println("Write error:", err)
			return
		}
	}
}