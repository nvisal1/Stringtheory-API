package main

import (
	env "github.com/joho/godotenv"
	"log"
	"net/http"
	"stringtheory/courses"
	"stringtheory/courses/adapters"
	"stringtheory/exercises"
	"stringtheory/game"
	guitar_interaction "stringtheory/guitar-interaction"
	"stringtheory/lessons"
	user_authentication "stringtheory/user-authentication"
	user_curriculum_progress "stringtheory/user-curriculum-progress"
	user_management "stringtheory/user-management"

	database "stringtheory/drivers"
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
	initializeServiceModules()
	openInternalServiceAdapters()
}

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

func openInternalServiceAdapters() {
	adapters.OpenInternalAdapter()
	exercises.OpenInternalAdapter()
	lessons.OpenInternalAdapter()
	user_management.OpenInternalAdapter()
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
