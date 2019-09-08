package user_management

import "Stringtheory-API/shared"

type dataStore interface {
	getUser(un string) (shared.User, error)
	createUser(u shared.User) error
}
