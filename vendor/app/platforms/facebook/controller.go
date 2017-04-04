package facebook

import "app/bot"
import "log"

func handleMessageEvent(message messaging, sender recipient) {
	log.Printf("Received Facebook message from user %s: %s", message.Sender.ID, message.Message.Text)
	bot.HandleMessage(message.Message.Text, message.Sender.ID, NewSender(sender))
}

func handlePostbackEvent(postback receivedPostback, sender recipient) {
}
