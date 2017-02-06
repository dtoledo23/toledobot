package facebook

import (
  "fmt"
)

func HandleMessageEvent(message Message) {
  fmt.Println("I received a message:", message.Text)
}

func HandlePostbackEvent(postback Postback) {
  fmt.Println("I received a message:", postback.Payload)
}
