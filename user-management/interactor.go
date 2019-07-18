package user_management

import "stringtheory/shared"

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
	u, err := sm.ds.getUser(un)
	return u, err
}