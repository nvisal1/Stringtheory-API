package user_management

import (
	"errors"
	"stringtheory/shared"
)

// register accepts an object nu of type newUser.
// The function checks to ensure that a user with
// the same username or email does not exist.
// If no error is thrown, the function will use
// the bcrypt package to hash the given password.
// The new user is saved with the hashed password.
//
// Finally, this function reaches out to the
// user-authentication module to retrieve
// a token for the new user.
func register(nu newUser) (userToken, error) {

}

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