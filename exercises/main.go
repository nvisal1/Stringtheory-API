package exercises

import (
	"Stringtheory-API/drivers/database"
	"Stringtheory-API/exercises/adapters/http"
	. "Stringtheory-API/exercises/drivers"
	. "Stringtheory-API/exercises/service-module"
)

// InitializeModule is an exported function.
//
// This function looks for SERVICE_ENVIRONMENT
// environment variable. It uses this variable to
// determine whether or not to load test stubs or
// production code.
func InitializeModule() {
	datastore := ModuleMongoDataStore{ database.GetConnection().Db}
	httpAdapter := http.ModuleHttpAdapter{}

	NewExercisesModule(httpAdapter, datastore)
	ExercisesModule.HttpAdapter.InitializeAdapter()
}

