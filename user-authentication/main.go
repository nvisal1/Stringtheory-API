package user_authentication

import (
	. "Stringtheory-API/user-authentication/adapters/http"
	. "Stringtheory-API/user-authentication/drivers/password-service"
	. "Stringtheory-API/user-authentication/drivers/service-communication"
	. "Stringtheory-API/user-authentication/drivers/token-generator"
	. "Stringtheory-API/user-authentication/service-module"
)

// InitializeModule is an exported function.
//
// This function looks for SERVICE_ENVIRONMENT
// environment variable. It uses this variable to
// determine whether or not to load test stubs or
// production code.
func InitializeModule() {
	passwordService := NewBcryptPasswordService()
	tokenGenerator := NewJWTGenerator()
	serviceCommunicator := NewServiceCommunicator()
	httpAdapter := NewHttpAdapter()

	NewUserAuthenticationModule(httpAdapter, passwordService, tokenGenerator, serviceCommunicator)
	UserAuthenticationModule.HttpAdapter.InitializeAdapter()
}