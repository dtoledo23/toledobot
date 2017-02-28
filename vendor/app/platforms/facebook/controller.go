package facebook

import (
	"app/bot"
	"fmt"
)

func handleMessageEvent(message receivedMessage, sender recipient) {
	fmt.Println("I received a message:", message.Text)
	bot.HandleMessage(message.Text, NewSender(sender))
}

func handlePostbackEvent(postback receivedPostback, sender recipient) {
	fmt.Println("I received a message:", postback.Payload)
}
