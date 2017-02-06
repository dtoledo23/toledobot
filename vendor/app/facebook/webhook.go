package facebook

import (
  "os"
  "fmt"
  "log"
  "encoding/json"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

// Facebook's code for validation adapted to Golang.
func WebhookValidation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  VERIFY_TOKEN := os.Getenv("FB_VERIFY_TOKEN")
  query := r.URL.Query()

  if (query.Get("hub.mode") == "subscribe") &&
    (query.Get("verify_token") == VERIFY_TOKEN) {
    log.Println("Validating webhook")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, query.Get("hub.challenge"))
  } else {
    log.Println("Failed validation. Make sure the validation tokens match.")
    w.WriteHeader(http.StatusForbidden)
  }
}

func EventHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var event FacebookEvent

  // Parse JSON
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&event)
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
        fmt.Println(event.Message, event.Message.Empty())
        // Received message
        if !event.Message.Empty() {
          HandleMessageEvent(event.Message)
        }

        // Received Postback
        if !event.Postback.Empty() {
          HandlePostbackEvent(event.Postback)
        }
      }
    }
  }
}
