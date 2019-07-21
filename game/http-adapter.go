package game

import "net/http"

func initializeAdapter() {
	http.Handle("/game/", http.StripPrefix("/game", http.FileServer(http.Dir("./game/app"))))
	http.HandleFunc("/game/play", phaser)
}

func phaser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}