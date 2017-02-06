package facebook

// ----- Structure of Facebook's request body for events -------
type FacebookEvent struct {
  Object string `json:"object"`
  Entry []Entry `json:"entry"`
}

type Entry struct {
  Id string  `json:"id"`
  Time int64 `json:"time"`
  Messaging []Messaging `json:"messaging"`
}

type Messaging struct {
  Recipient struct {
    Id string `json:"id"`
  } `json:"recipient"`

  Sender struct {
    Id string `json:"id"`
  } `json:"sender"`

  Timestamp int64 `json:"timestamp"`

  // Optional depending on event type.
	Message Message `json:"message,omitempty"`
  Postback Postback `json:"postback,omitempty"`
}

// -------- Currently supported events. --------
// Message event
type Message struct {
  Mid string `json:"mid"`
  Seq int64 `json:"seq"`
  Text string `json:"text"`
  Attachments []Attachment `json:"attachments"`
  QuickReply QuickReply `json:"quick_reply"`
}

func (m Message) Empty() bool {
    return (m.Mid == "")
}

// Postback event
type Postback struct {
  Payload string `json:"payload"`
}

func (p Postback) Empty() bool {
    return (p.Payload == "")
}

// ----------- Complementary object types -----------
type Attachment struct {
  Type string `json:"type"`
  Payload Multimedia `json:"payload"`
}

type Multimedia struct {
  Url string `json:"url"`
}

type QuickReply struct {
  Payload string `json:"payload"`
}
