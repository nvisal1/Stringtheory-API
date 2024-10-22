package user_management

import (
	database2 "Stringtheory-API/drivers/database"
	"os"
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
			ia: userManagementAdapter{},
			ds: moduleMongoDataStore {
				database2.GetConnection().Db,
			},
			tds: stubMongoDataStore{},
			se: se,
		}
		sm.ha.InitializeAdapter()
		sm.ia.InitializeAdapter()
	}
}