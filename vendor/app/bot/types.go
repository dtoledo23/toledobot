package bot

type botAnswer struct {
	ID       string            `json:"_id"`
	Type     string            `json:"type"`
	Payload  string            `json:"payload"`
	Entities map[string]string `json:"entities"`
}

type botSession struct {
	ID      string            `json:"_id"`
	UserID  string            `json:"userId"`
	Context map[string]string `json:"context"`
}
