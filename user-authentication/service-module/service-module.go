package service_module

import (
	"Stringtheory-API/shared"
	"Stringtheory-API/user-authentication/interfaces"
)

var UserAuthenticationModule *ServiceModule

type ServiceModule struct {
	HttpAdapter shared.Adapter
	PasswordService interfaces.PasswordService
	TokenGenerator interfaces.TokenGenerator
	ServiceCommunicator interfaces.ServiceCommunicator
}

func NewUserAuthenticationModule(
	httpAdapter shared.Adapter,
	passwordService interfaces.PasswordService,
	tokenGenerator interfaces.TokenGenerator,
	serviceCommunicator interfaces.ServiceCommunicator,
) {
	UserAuthenticationModule = &ServiceModule{
		HttpAdapter:         httpAdapter,
		PasswordService:     passwordService,
		TokenGenerator:      tokenGenerator,
		ServiceCommunicator: serviceCommunicator,
	}
}


