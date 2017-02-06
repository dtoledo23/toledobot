package facebook

type SentPayload struct {
  Recipient Recipient `json:"recipient"`
  Message SentMessage `json:"message"`
}

type SentMessage struct {
  Text string `json:"text"`
}
