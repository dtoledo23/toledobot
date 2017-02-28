package main

import (
	"app/routes"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Run server.
	// Address = HOST:PORT
	address := strings.Join([]string{os.Getenv("HOST"), os.Getenv("PORT")}, ":")
	log.Println("ToledoBot running on ", address)
	log.Fatal(http.ListenAndServe(address, routes.Router()))

}
