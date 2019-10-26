package user_curriculum_progress

import (
	"Stringtheory-API/drivers/database"
	. "Stringtheory-API/user-curriculum-progress/adapter/http"
	. "Stringtheory-API/user-curriculum-progress/adapter/queue"
	. "Stringtheory-API/user-curriculum-progress/drivers/data-store"
	. "Stringtheory-API/user-curriculum-progress/drivers/service-communication"
	. "Stringtheory-API/user-curriculum-progress/service-module"
)

// InitializeModule is an exported function.
//
// This function looks for SERVICE_ENVIRONMENT
// environment variable. It uses this variable to
// determine whether or not to load test stubs or
// production code.
func InitializeModule() {
	httpAdapter := NewHttpAdapter()
	queueAdapter := NewQueueAdapter(50)
	dataStore := NewDatastore(database.GetConnection().Db)
	messageStore := NewMessageStore()
	serviceCommunicator := NewServiceCommunicator()

	NewUserCurriculumProgressModule(httpAdapter, queueAdapter, dataStore, messageStore, serviceCommunicator)
	UserCurriculumProgressModule.HttpAdapter.InitializeAdapter()
	UserCurriculumProgressModule.QueueAdapter.InitializeAdapter()
}