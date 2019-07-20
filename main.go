package main

import (
	"github.com/gorilla/websocket"
	env "github.com/joho/godotenv"
	"log"
	"net/http"
	"stringtheory/courses"
	user_management "stringtheory/user-management"

	database "stringtheory/drivers"
	"stringtheory/user-authentication"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)


type Message struct {
	Note string
}


// init is responsible for creating all of the necessary
// components to run the service.
//
// This function starts by loading all of the
// provided environment variables
//
// After the environment variables are loaded, this function
// uses the service database connection
// interface to establish a single connection.
//
// After the connection is established, this function initializes
// all of its required service modules.
//
// The order of these operations are important because the
// service modules ask the database connection interface for
// for the connection. If the connection does not exist, then
// the service will exit. The environment variables must be loaded first
// because they are used when establishing a database connection.
func init() {
	env.Load()
	database.Build()
	courses.InitializeModule()
	user_authentication.InitializeModule()
	user_management.InitializeModule()
	user_management.OpenInternalAdapter()
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