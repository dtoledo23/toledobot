package routes

import (
	"app/platforms/facebook"
	"app/platforms/sockets"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Router defines API routes
func Router() *httprouter.Router {
	router := httprouter.New()

	// Root
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hi! I am ToledoBot up and running\n")
	})

	// Facebook webhook
	router.GET("/facebook/webhook", facebook.WebhookValidation)
	router.POST("/facebook/webhook", facebook.EventHandler)

	// Sockets webhook
	router.GET("/ws/webhook", sockets.ServeWs)

	return router
}
