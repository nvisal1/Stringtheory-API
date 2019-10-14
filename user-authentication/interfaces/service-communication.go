package interfaces

import "Stringtheory-API/shared"

type ServiceCommunicator interface {
	GetUser(username string) (shared.User, error)
	CreateUser(user shared.User) error
}