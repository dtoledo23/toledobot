package bot

import (
	"app/platforms"
	"log"
)

const (
	witAPI = "api.wit.ai"
)

type answer struct {
	Type     string            `json:"type"`
	Payload  string            `json:"payload"`
	Entities map[string]string `json:"type"`
}

// HandleMessage defines what to do when receiving a message.
func HandleMessage(message string, sender platforms.Sender) {
	log.Println("Received:", message)
	var response string

	switch entities := getEntities(message); entities["intent"] {
	case "greetings":
		response = "Hola! Soy TrokoBot"

	case "address":
		response = "Paseo del Hospicio #22, San Juan de Dios 44360 Guadalajara local 2151"

	case "schedule":
		response = "Lunes a Sabado de 9:00am a 7:00pm"

	case "thanks":
		response = "No te preocupes"

	case "bye":
		response = "adios! :)"

	case "ok":
		response = "Chido"

	case "insult":
		response = "Oye tranquilo viejo"

	case "memo":
		response = "Pongase a programar 😎"

	default:
		response = "Lo siento. No pude entender."
	}

	sender.SendText(response)
}
