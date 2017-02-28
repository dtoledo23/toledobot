package facebook

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// WebhookValidation Facebook's code for validation adapted to Golang.
func WebhookValidation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	verifyToken := os.Getenv("FB_VERIFY_TOKEN")
	query := r.URL.Query()
	if (query.Get("hub.mode") == "subscribe") &&
		(query.Get("hub.verify_token") == verifyToken) {
		log.Println("Validating webhook")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, query.Get("hub.challenge"))
	} else {
		log.Println("Failed validation. Make sure the validation tokens match.")
		w.WriteHeader(http.StatusForbidden)
	}
}

// EventHandler handles post request from Facebook for incomming messages
func EventHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var event receivedPayload

	// Parse JSON
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		// Respond to Facebook no matter what.
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
		panic(err)
	}

	// Cached only page related events.
	if event.Object == "page" {
		for _, entry := range event.Entry {
			// Not used now. Included for reference.
			// pageID := entry.Id
			// eventTimestamp := entry.Time

			// Handle all events received.
			for _, event := range entry.Messaging {
				// Received message
				if !event.Message.Empty() {
					handleMessageEvent(event.Message, event.Sender)
				}

				// Received Postback
				if !event.Postback.Empty() {
					handlePostbackEvent(event.Postback, event.Sender)
				}
			}
		}
	}
}
