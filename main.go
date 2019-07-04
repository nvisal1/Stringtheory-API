package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)


type Message struct {
	Note string
}

func main() {
	hub := newHub()
	go hub.run()
	http.Handle("/phaser/", http.StripPrefix("/phaser", http.FileServer(http.Dir("./phaser/app"))))
	http.HandleFunc("/phaser/play", phaser)
	http.HandleFunc("/record", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func phaser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}