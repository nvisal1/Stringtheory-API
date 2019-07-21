package user_authentication

import (
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
			moduleHttpAdapter{},
			se,
		}
		sm.ha.InitializeAdapter()
	}
}