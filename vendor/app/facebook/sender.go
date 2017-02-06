package facebook

import (
  "os"
  "log"
  "time"
  "bytes"
  "strings"
  "net/http"
  "encoding/json"
)

var MESSENGER_API =  "https://graph.facebook.com/v2.6/me/messages"

func SendText(recipient Recipient, text string) {
  message := SentMessage{text}
  send(recipient, message)
}

func send(recipient Recipient, message SentMessage) {
  payload := SentPayload{
    Recipient: recipient,
    Message: message,
  }

  callSendApi(payload)
}

func callSendApi(payload SentPayload) {
  // Parse payload into JSON.
  buffer := new(bytes.Buffer)
  json.NewEncoder(buffer).Encode(payload)

  // Setup POST request.
  url := strings.Join([]string{MESSENGER_API, "?access_token=", os.Getenv("FB_PAGE_TOKEN")}, "")
  client := &http.Client{Timeout: time.Second * 10}
  resp, err := client.Post(url, "application/json; charset=utf-8", buffer)


  if err != nil {
    log.Panic(err)
  }

  defer resp.Body.Close()
}
