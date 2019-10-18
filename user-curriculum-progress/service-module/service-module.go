package service_module

import (
	"Stringtheory-API/shared"
	"Stringtheory-API/user-curriculum-progress/interfaces"
)

var UserCurriculumProgressModule *ServiceModule

type ServiceModule struct {
	HttpAdapter  shared.Adapter
	Datastore interfaces.Datastore
	ServiceCommunicator interfaces.ServiceCommunicator
}

func NewUserCurriculumProgressModule(
	httpAdapter shared.Adapter,
	datastore interfaces.Datastore,
	serviceCommunicator interfaces.ServiceCommunicator,
) {
	UserCurriculumProgressModule = &ServiceModule{
		httpAdapter:         httpAdapter,
		Datastore:           datastore,
		ServiceCommunicator: serviceCommunicator,
	}
}
