package user_management

import "stringtheory/shared"

type dataStore interface {
	getUser(un string) (shared.User, error)
}

type httpAdapter interface {
	initializeAdapter()
}