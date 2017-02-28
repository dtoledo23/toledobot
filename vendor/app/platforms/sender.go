package platforms

type Sender interface {
	SendText(message string) error
	SendImage(url string) error
}
