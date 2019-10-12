package service_module

import (
	localInterfaces "Stringtheory-API/courses/interfaces"
	sharedTypes "Stringtheory-API/shared"
	"errors"
	"os"
)

var CoursesModule ServiceModule

type ServiceModule struct {
	HttpAdapter sharedTypes.Adapter
	Datastore localInterfaces.Datastore
}

func (module ServiceModule) ResetDatastore(datastore localInterfaces.Datastore) error {
	serviceEnvironment, exists := os.LookupEnv("SERVICE_ENVIRONMENT")
	if exists {
		if serviceEnvironment == "test" {
			module.Datastore = datastore
			return nil
		}
		return errors.New("the service environment variable is not set")
	}
	return errors.New("a module's datastore can only be reset when the service environment is set to test")
}
