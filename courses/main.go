package courses

import (
	"Stringtheory-API/courses/adapters/http"
	"Stringtheory-API/courses/drivers"
	localInterfaces "Stringtheory-API/courses/interfaces"
	"Stringtheory-API/courses/service-module"
	"Stringtheory-API/drivers/database"
	"errors"
	"log"
	"os"
)

// InitializeModule is an exported function.
//
// This function looks for SERVICE_ENVIRONMENT
// environment variable. It uses this variable to
// determine whether or not to load test stubs or
// production code.
func InitializeModule() {
	datastore, err := selectDatastore()
	if err != nil {
		log.Fatal(err)
	}
	service_module.CoursesModule = service_module.ServiceModule{
		http.ModuleHttpAdapter{},
		datastore,
	}
	service_module.CoursesModule.HttpAdapter.InitializeAdapter()
}

func selectDatastore() (localInterfaces.Datastore, error) {
	serviceEnvironment, exists := os.LookupEnv("SERVICE_ENVIRONMENT")
	if exists {
		if serviceEnvironment == "test" {
			return drivers.StubMongoDataStore{} ,nil
		} else {
			return drivers.ModuleMongoDataStore{
				database.GetConnection().Db,
			}, nil
		}
		return drivers.StubMongoDataStore{}, errors.New("the service environment variable is not set")
	}
	return drivers.StubMongoDataStore{}, errors.New("a module's datastore can only be reset when the service environment is set to test")
}