package sockets

import (
	"github.com/gorilla/websocket"
)

// Client defines method to send messages to socket client.
type Client struct {
	conn *websocket.Conn
}

// NewClient creates a new sender object for socket client.
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

// SendText sends a text message to a socket client.
func (client *Client) SendText(text string) error {
	return client.conn.WriteMessage(websocket.TextMessage, []byte(text))
}

// SendImage sends a image message to a socket client.
func (client *Client) SendImage(url string) error {
	// TODO(dtoledo23): Implement logic for sending images.
	return nil
}
