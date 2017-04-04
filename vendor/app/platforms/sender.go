package platforms

// Sender defines an interface for supporting answering to messages in various platforms.
type Sender interface {
	SendText(message string) error
	SendImage(url string) error
}
