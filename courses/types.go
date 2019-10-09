package courses

import (
	localInterfaces "Stringtheory-API/courses/interfaces"
	sharedTypes "Stringtheory-API/shared"
	"errors"
	"os"
)

type serviceModule struct {
	httpAdapter sharedTypes.Adapter
	internalAdapter sharedTypes.Adapter
	datastore localInterfaces.Datastore
}

func (module serviceModule) ResetDatastore(datastore localInterfaces.Datastore) error {
	serviceEnvironment, exists := os.LookupEnv("SERVICE_ENVIRONMENT")
	if exists {
		if serviceEnvironment == "test" {
			module.datastore = datastore
			return nil
		}
		return errors.New("the service environment variable is not set")
	}
	return errors.New("a module's datastore can only be reset when the service environment is set to test")
}