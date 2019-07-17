package user_authentication

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
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
	u := sm.DataStore().getUser(lc.username)
	bytePass := []byte(lc.password)
	if comparePasswords(u.password, bytePass) {



		return userToken{
			tokenString,
		}, nil
	}
	return userToken{
		nil,
	}, errors.New("Password is incorrect")

}

