package facebook

// ----------- Complementary object types -----------
type recipient struct {
	ID string `json:"id"`
}

type attachment struct {
	Type    string     `json:"type"`
	Payload multimedia `json:"payload"`
}

type multimedia struct {
	URL string `json:"url"`
}

type quickReply struct {
	Payload string `json:"payload"`
}
