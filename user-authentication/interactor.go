package user_authentication

import (
	"errors"
	"Stringtheory-API/shared"
)

// generateToken accepts an object u of type LoginCredentials.
// The function uses the username of the given LoginCredentials
// variable to fetch the appropriate user from the database.
//
// If a user is returned, the function will use the bcrypt
// package to compare the hashed password  with the plaintext
// password.
//
// If the passwords match, the function will construct
// an object of type User. The jwt-go package is
// then used to create a token. The User object is added
// as a claim.
//
// If no errors are returned during this process, the function
// returns an object of type UserToken
func processLogin(lc loginCredentials) (userToken, error) {
	u, err := getUser(lc.Username)
	if err != nil {
		return userToken{
			"",
		}, errors.New("Username is incorrect")
	}
	if comparePasswords(u.Password, lc.Password) {
		su := shared.SecureUser{
			Username: u.Username,
			Email: u.Email,
			Name: u.Name,
		}
		token, err := generateToken(su)
		if err != nil {
			return userToken{
				"",
			}, err
		}

		return userToken{
			token,
		}, nil
	}
	return userToken{
		"",
	}, errors.New("Password is incorrect")
}

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
func processRegistration(u shared.User) (userToken, error) {
	_, err := getUser(u.Username)
	// If err, then user does not exist for given username
	if err != nil {
		hp, err := encryptPassword(u.Password)
		if err != nil {
			return userToken{
				"",
			}, err
		}
		u.Password = hp
		err = createUser(u)
		if err != nil {
			return userToken{
				"",
			}, err
		}
		su := shared.SecureUser{
			Username: u.Username,
			Email: u.Email,
			Name: u.Name,
		}
		token, err := generateToken(su)
		if err != nil {
			return userToken{
				"",
			}, err
		}
		return userToken{
			token,
		}, nil
	}
	return userToken{
		"",
	}, errors.New("A user with this username already exists")
}
