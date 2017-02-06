package facebook

import (
  "fmt"
)

func HandleMessageEvent(message ReceivedMessage, sender Recipient) {
  fmt.Println("I received a message:", message.Text)
  SendText(sender, message.Text)
}

func HandlePostbackEvent(postback ReceivedPostback, sender Recipient) {
  fmt.Println("I received a message:", postback.Payload)
}
