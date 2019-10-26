package service_module

import (
	"Stringtheory-API/shared"
	"Stringtheory-API/user-curriculum-progress/interfaces"
)

var UserCurriculumProgressModule *ServiceModule

type ServiceModule struct {
	HttpAdapter  shared.Adapter
	QueueAdapter shared.Adapter
	DataStore interfaces.DataStore
	MessageStore interfaces.MessageStore
	ServiceCommunicator interfaces.ServiceCommunicator
}

func NewUserCurriculumProgressModule(
	httpAdapter shared.Adapter,
	queueAdapter shared.Adapter,
	dataStore interfaces.DataStore,
	messageStore interfaces.MessageStore,
	serviceCommunicator interfaces.ServiceCommunicator,
) {
	UserCurriculumProgressModule = &ServiceModule{
		HttpAdapter:         httpAdapter,
		QueueAdapter:        queueAdapter,
		DataStore:           dataStore,
		MessageStore:        messageStore,
		ServiceCommunicator: serviceCommunicator,
	}
}
