package main

import (
  "fmt"
  "log"
  "app/routes"
  "net/http"
)

func main() {

    // Run server.
    log.Fatal(http.ListenAndServe(":8080", routes.Router()))
}
