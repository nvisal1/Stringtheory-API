package user_management

import (
	"errors"
	"Stringtheory-API/shared"
)



func getUser(un string) (shared.User, error) {
	u, err := sm.DataStore().getUser(un)
	return u, err
}

func createUser(u shared.User) error {
	err := sm.DataStore().createUser(u)
	if err != nil {
		return errors.New("User could not be created")
	}
	return nil
}