package user_authentication

import (
	"errors"
	"stringtheory/shared"
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
	u, err := getUser(lc.username)
	if err != nil {
		return userToken{
			nil,
		}, errors.New("Username is incorrect")
	}
	if comparePasswords(u.Password, lc.password) {
		su := shared.SecureUser{
			Username: u.Username,
			Email: u.Email,
			Name: u.Name,
		}
		token, err := generateToken(su)
		if err != nil {
			return userToken{
				nil,
			}, err
		}

		return userToken{
			token,
		}, nil
	}
	return userToken{
		nil,
	}, errors.New("Password is incorrect")
}


func processRegistration(nu newUser) (userToken, error) {
	_, err := getUser(nu.username)
	// If err, then user does not exist for given username
	if err != nil {
		hp, err := encryptPassword(nu.password)
		if err != nil {
			return userToken{
				nil,
			}, err
		}
		nu.password = hp

	}
	return userToken{
		nil,
	}, errors.New("A user with this username already exists")
}
