package login

import (
	"Stringtheory-API/shared"
	. "Stringtheory-API/user-authentication/service-module"
	. "Stringtheory-API/user-authentication/types"
	"errors"
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
func Login(loginCredentials LoginCredentials) (UserToken, error) {
	user, err := UserAuthenticationModule.ServiceCommunicator.GetUser(loginCredentials.Username)
	if err != nil {
		return UserToken{}, errors.New("the provided username is incorrect")
	}
	if UserAuthenticationModule.PasswordService.ComparePasswords(user.Password, loginCredentials.Password) {
		secureUser := shared.SecureUser{
			Username: user.Username,
			Email: user.Email,
			Name: user.Name,
		}
		token, err := UserAuthenticationModule.TokenGenerator.GenerateToken(secureUser)
		if err != nil {
			return UserToken{}, err
		}

		return UserToken{Token: token}, nil
	}
	return UserToken{}, errors.New("the provided password is incorrect")
}