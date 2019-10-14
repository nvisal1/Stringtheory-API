package lessons

import (
	"Stringtheory-API/drivers/database"
	. "Stringtheory-API/lessons/adapters/http"
	. "Stringtheory-API/lessons/drivers/data-store"
	. "Stringtheory-API/lessons/drivers/service-communication"
	. "Stringtheory-API/lessons/service-module"
)

// InitializeModule is an exported function.
//
// This function looks for SERVICE_ENVIRONMENT
// environment variable. It uses this variable to
// determine whether or not to load test stubs or
// production code.
func InitializeModule() {
	datastore := ModuleMongoDataStore{database.GetConnection().Db}
	httpAdapter := ModuleHttpAdapter{}
    serviceCommunicator := ServiceCommunicator{}

	NewLessonsModule(httpAdapter, datastore, serviceCommunicator)
    LessonsModule.HttpAdapter.InitializeAdapter()
}

