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
func generateToken(lc loginCredentials) (userToken, error) {
	key, exists := os.LookupEnv("KEY")
	if exists {
		u := sm.DataStore().getUser(lc.username)
		bytePass := []byte(lc.password)
		if comparePasswords(u.password, bytePass) {
			tokenUser := secureUser{
				username: u.username,
				name:     u.name,
				email:    u.email,
			}
			token := jwt.New(jwt.SigningMethodHS256)
			token.Claims = jwt.MapClaims{
				"exp":  time.Now().Add(time.Hour * 72).Unix(),
				"iat":  time.Now().Unix(),
				"user": tokenUser,
			}
			byteKey := []byte(key)
			tokenString, err := token.SignedString(byteKey)
			if err != nil {
				return userToken{
					token: nil,
				}, err
			}
			return userToken{
				tokenString,
			}, nil
		}
		return userToken{
			nil,
		}, errors.New("Password is incorrect")
	}
	return userToken{
		nil,
	}, errors.New("Key is not set correctly")
}


// comparePasswords is responsible for comparing
// the user's plaintext password to
// the hashed password
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
