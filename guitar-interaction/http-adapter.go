package guitar_interaction

import (
	"net/http"
	"stringtheory/shared"
)

func initializeAdapter() {
	hub := newHub()
	go hub.run()
	http.HandleFunc("/play", shared.Authenticate()(func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	}))
}

