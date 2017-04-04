package sockets

type socketsMessage struct {
	Payload   string `json:"payload"`
	Timestamp int64  `json:"type"`
	UserID    string `json:"userId"`
}
