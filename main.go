package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Note string
}

func main() {
	http.Handle("/phaser/", http.StripPrefix("/phaser", http.FileServer(http.Dir("./phaser/app"))))
	http.HandleFunc("/phaser/play", phaser)
	http.HandleFunc("/record", handleConnections)
	go handleMessages()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func phaser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error")
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	msg := <- broadcast
	log.Print(msg)
	for client := range clients {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("err")
			client.Close()
			delete(clients, client)
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}