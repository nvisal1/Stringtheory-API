package guitar_interaction

import (
	"Stringtheory-API/drivers/router"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func initializeAdapter() {
	hub := newHub()
	go hub.run()
	router.GetRouter().GET("/play", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		serveWs(hub, w, r)
	})
}

