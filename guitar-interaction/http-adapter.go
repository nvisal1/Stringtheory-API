package guitar_interaction

import (
	"net/http"
)

func initializeAdapter() {
	hub := newHub()
	go hub.run()
	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
}

