package facebook

// ----- Structure of Facebook's request body for events -------
type ReceivedPayload struct {
  Object string `json:"object"`
  Entry []Entry `json:"entry"`
}

type Entry struct {
  Id string  `json:"id"`
  Time int64 `json:"time"`
  Messaging []Messaging `json:"messaging"`
}

type Messaging struct {
  Recipient Recipient `json:"recipient"`
  Sender Recipient `json:"sender"`
  Timestamp int64 `json:"timestamp"`

  // Optional depending on event type.
	Message ReceivedMessage `json:"message,omitempty"`
  Postback ReceivedPostback `json:"postback,omitempty"`
}

// -------- Currently supported events. --------
// Message event
type ReceivedMessage struct {
  Mid string `json:"mid"`
  Seq int64 `json:"seq"`
  Text string `json:"text"`
  Attachments []Attachment `json:"attachments"`
  QuickReply QuickReply `json:"quick_reply"`
}

func (m ReceivedMessage) Empty() bool {
    return (m.Mid == "")
}

// Postback event
type ReceivedPostback struct {
  Payload string `json:"payload"`
}

func (p ReceivedPostback) Empty() bool {
    return (p.Payload == "")
}
