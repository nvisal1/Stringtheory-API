package service_module

import (
	"Stringtheory-API/exercises/interfaces"
	sharedTypes "Stringtheory-API/shared"
)

var ExercisesModule *ServiceModule

type ServiceModule struct {
	HttpAdapter sharedTypes.Adapter
	Datastore interfaces.Datastore
}

func NewExercisesModule(httpAdapter sharedTypes.Adapter, datastore interfaces.Datastore) {
	ExercisesModule = &ServiceModule{
		HttpAdapter: httpAdapter,
		Datastore:   datastore,
	}
}


