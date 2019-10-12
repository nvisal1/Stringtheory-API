package courses

import (
	"Stringtheory-API/courses/adapters/http"
	. "Stringtheory-API/courses/drivers"
	. "Stringtheory-API/courses/service-module"
	"Stringtheory-API/drivers/database"
)

// InitializeModule is an exported function.
//
// This function looks for SERVICE_ENVIRONMENT
// environment variable. It uses this variable to
// determine whether or not to load test stubs or
// production code.
func InitializeModule() {
	datastore := ModuleMongoDataStore{database.GetConnection().Db}
	httpAdapter := http.ModuleHttpAdapter{}

	NewCoursesModule(httpAdapter, datastore)
	CoursesModule.HttpAdapter.InitializeAdapter()
}

