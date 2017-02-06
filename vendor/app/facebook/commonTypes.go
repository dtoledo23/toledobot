package facebook

// ----------- Complementary object types -----------
type Recipient struct {
  Id string `json:"id"`
}

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
