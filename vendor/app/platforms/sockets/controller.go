package sockets

import (
	"fmt"

	"app/bot"

	"github.com/gorilla/websocket"
)

func handleSocket(conn *websocket.Conn) error {
	client := &Client{conn}
	defer client.conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		if messageType != websocket.TextMessage {
			return fmt.Errorf("%s", "Message type not supported")
		}

		bot.HandleMessage(string(message[:]), client)
	}
}

func handleSocketIO(conn *websocket.Conn) error {
	client := &Client{conn}
	defer client.conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		if messageType != websocket.TextMessage {
			return fmt.Errorf("%s", "Message type not supported")
		}

		bot.HandleMessage(string(message[:]), client)
	}
}
