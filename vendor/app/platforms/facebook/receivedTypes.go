package facebook

// ----- Structure of Facebook's request body for events -------
type receivedPayload struct {
	Object string  `json:"object"`
	Entry  []entry `json:"entry"`
}

type entry struct {
	ID        string      `json:"id"`
	Time      int64       `json:"time"`
	Messaging []messaging `json:"messaging"`
}

type messaging struct {
	Recipient recipient `json:"recipient"`
	Sender    recipient `json:"sender"`
	Timestamp int64     `json:"timestamp"`

	// Optional depending on event type.
	Message  receivedMessage  `json:"message,omitempty"`
	Postback receivedPostback `json:"postback,omitempty"`
}

// -------- Currently supported events. --------
// Message event
type receivedMessage struct {
	Mid         string       `json:"mid"`
	Seq         int64        `json:"seq"`
	Text        string       `json:"text"`
	Attachments []attachment `json:"attachments"`
	QuickReply  quickReply   `json:"quick_reply"`
}

func (m receivedMessage) Empty() bool {
	return (m.Mid == "")
}

// Postback event
type receivedPostback struct {
	Payload string `json:"payload"`
}

func (p receivedPostback) Empty() bool {
	return (p.Payload == "")
}
