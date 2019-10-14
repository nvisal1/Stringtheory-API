package register

import (
	"Stringtheory-API/shared"
	. "Stringtheory-API/user-authentication/service-module"
	. "Stringtheory-API/user-authentication/types"
	"errors"
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
func Register(user shared.User) (UserToken, error) {
	_, err := UserAuthenticationModule.ServiceCommunicator.GetUser(user.Username)
	// If err, then user does not exist for given username
	if err != nil {
		hashedPassword, err := UserAuthenticationModule.PasswordService.EncryptPassword(user.Password)
		if err != nil {
			return UserToken{}, err
		}
		user.Password = hashedPassword
		err = UserAuthenticationModule.ServiceCommunicator.CreateUser(user)
		if err != nil {
			return UserToken{}, err
		}
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
	return UserToken{}, errors.New("A user with this username already exists")
}