package bot

import (
	"app/platforms"
	"fmt"
)

// HandleMessage defines what to do when receiving a message.
func HandleMessage(message string, sender platforms.Sender) {
	fmt.Println("Ill do my best to answer to ", message)
	sender.SendText("Hi")
}
