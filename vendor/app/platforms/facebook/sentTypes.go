package facebook

type sentPayload struct {
	Recipient recipient   `json:"recipient"`
	Message   sentMessage `json:"message"`
}

type sentMessage struct {
	Text string `json:"text"`
}
