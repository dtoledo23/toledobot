package sockets

import (
	"app/bot"
	"fmt"
	"log"

	"encoding/json"

	"github.com/gorilla/websocket"
)

func handleSocket(conn *websocket.Conn) error {
	client := &Client{conn}
	defer client.conn.Close()

	for {
		messageType, rawMessage, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		if messageType != websocket.TextMessage {
			return fmt.Errorf("%s", "Message type not supported")
		}

		var message socketsMessage
		json.Unmarshal(rawMessage, &message)

		log.Printf("Received WebSocket message from user %s: %s", message.UserID, message.Payload)
		bot.HandleMessage(message.Payload, message.UserID, client)
	}
}
