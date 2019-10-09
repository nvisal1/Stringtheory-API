package courses

import (
	"Stringtheory-API/courses/adapters"
	"Stringtheory-API/courses/adapters/http"
	"Stringtheory-API/courses/drivers"
	database "Stringtheory-API/drivers/database"
	"log"
)

var CoursesModule serviceModule

// InitializeModule is an exported function.
//
// This function looks for SERVICE_ENVIRONMENT
// environment variable. It uses this variable to
// determine whether or not to load test stubs or
// production code.
func InitializeModule() {
	datastore, err := createModuleDataStore()
	if err != nil {
		log.Fatal(err)
	}
	CoursesModule = serviceModule{
		http.ModuleHttpAdapter{},
		adapters.InternalAdapter{},
		datastore,
	}
	CoursesModule.httpAdapter.InitializeAdapter()
	CoursesModule.internalAdapter.InitializeAdapter()
}

func createModuleDataStore() (Datastore, error) {
	datastore := drivers.ModuleMongoDataStore{
			database.GetConnection().Db,
	}
	return datastore, nil
}
