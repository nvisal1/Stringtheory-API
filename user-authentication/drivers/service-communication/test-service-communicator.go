package service_communication

import "Stringtheory-API/shared"

type StubServiceCommunicator struct {}

func NewStubServiceCommunicator() *StubServiceCommunicator {
	return &StubServiceCommunicator{}
}

func (stubServiceCommunicator StubServiceCommunicator) GetUser(username string) (shared.User, error) {
	user := shared.User{
		Username: "test_username",
		Name:     "test_name",
		Email:    "test_email",
		Password: "test_password",
	}

	return user, nil
}

func (stubServiceCommunicator StubServiceCommunicator) CreateUser(user shared.User) error {
	return nil
}