package facebook

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var messangerAPI = "https://graph.facebook.com/v2.6/me/messages"

// Sender defines method to send messages to Facebook Messenger.
type Sender struct {
	recipient recipient
}

// NewSender creates a new sender object for Facebook.
func NewSender(recipient recipient) *Sender {
	return &Sender{
		recipient: recipient,
	}
}

// SendText sends a text message to Facebook Messenger.
func (fbs *Sender) SendText(text string) error {
	// TODO: Handle errors in correct way.
	message := sentMessage{text}
	send(fbs.recipient, message)
	return nil
}

// SendImage sends a image message to Facebook Messenger.
func (fbs *Sender) SendImage(url string) error {
	// TODO(dtoledo23): Implement logic for sending images.
	return nil
}

func send(recipient recipient, message sentMessage) {
	payload := sentPayload{
		Recipient: recipient,
		Message:   message,
	}

	callSendAPI(payload)
}

func callSendAPI(payload sentPayload) {
	// Parse payload into JSON.
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(payload)

	// Setup POST request.
	url := strings.Join([]string{messangerAPI, "?access_token=", os.Getenv("FB_PAGE_TOKEN")}, "")
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Post(url, "application/json; charset=utf-8", buffer)

	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()
}
