package service_communication

import (
	"Stringtheory-API/shared"
	user_management "Stringtheory-API/user-management"
)

type ServiceCommunicator struct {}

func NewServiceCommunicator() *ServiceCommunicator {
	return &ServiceCommunicator{}
}

func (ServiceCommunicator ServiceCommunicator) GetUser(username string) (shared.User, error) {
	user, err := user_management.GetInstance().GetUser(username)
	return user, err
}

func (serviceCommunicator ServiceCommunicator) CreateUser(user shared.User) error {
	err := user_management.GetInstance().CreateUser(user)
	return err
}