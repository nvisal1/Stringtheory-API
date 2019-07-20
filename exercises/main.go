package exercises

import (
	"os"
	database "stringtheory/drivers"
)

var sm serviceModule

// InitializeModule is an exported function.
//
// This function looks for SERVICE_ENVIRONMENT
// environment variable. It uses this variable to
// determine whether or not to load test stubs or
// production code.
func InitializeModule() {
	se, exists := os.LookupEnv("SERVICE_ENVIRONMENT")
	if exists {
		sm = serviceModule{
			ha: moduleHttpAdapter{},
			ds: moduleMongoDataStore {
				database.GetConnection().Db,
			},
			tds: stubMongoDataStore{},
			se: se,
		}
		sm.ha.initializeAdapter()
	}
}
