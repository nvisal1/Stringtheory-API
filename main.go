package main

import (
	"Stringtheory-API/courses"
	"Stringtheory-API/drivers/database"
	"Stringtheory-API/drivers/router"
	"Stringtheory-API/exercises"
	"Stringtheory-API/game"
	guitar_interaction "Stringtheory-API/guitar-interaction"
	"Stringtheory-API/lessons"
	user_authentication "Stringtheory-API/user-authentication"
	user_curriculum_progress "Stringtheory-API/user-curriculum-progress"
	user_management "Stringtheory-API/user-management"
	env "github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"io"
	"log"
	"net/http"
)

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
// all required service modules
//
// The order of these operations is important because the
// service modules ask the database connection interface for
// for the connection. If the connection does not exist, then
// the service will exit. The environment variables must be loaded first
// because they are used when establishing a database connection.
func init() {
	env.Load()
	database.Build()
	router.Build()
	initializeServiceModules()
}

// initializeServiceModules calls the InitializeModule function
// for every required service module. Module functionality
// will not be available unless its InitializeModule function
// is called.
func initializeServiceModules() {
	courses.InitializeModule()
	exercises.InitializeModule()
	game.InitializeModule()
	guitar_interaction.InitializeModule()
	lessons.InitializeModule()
	user_authentication.InitializeModule()
	user_curriculum_progress.InitializeModule()
	user_management.InitializeModule()
}

// main calls http.ListenAndServe to start the API.
// The service will exit if an error occurs
func main() {
	r := router.GetRouter()
	r.GET("/", index)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug: true,
	})
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":5000", handler))
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(w, "Welcome to the Stringtheory API!\n")
}
