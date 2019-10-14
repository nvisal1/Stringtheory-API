package service_module

import (
	. "Stringtheory-API/lessons/interfaces"
	"Stringtheory-API/shared"
)

var LessonsModule *ServiceModule

type ServiceModule struct {
	HttpAdapter  shared.Adapter
	Datastore  Datastore
	ServiceCommunicator ServiceCommunicator
}

func NewLessonsModule(httpAdapter shared.Adapter, datastore Datastore, serviceCommunicator ServiceCommunicator) {
	LessonsModule = &ServiceModule{
		HttpAdapter:         httpAdapter,
		Datastore:           datastore,
		ServiceCommunicator: serviceCommunicator,
	}
}