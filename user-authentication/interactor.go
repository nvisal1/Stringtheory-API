package user_authentication

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
// then used to create a token and add the User object
// as a claim.
//
// If no errors are returned during this process, the function
// returns an object of type UserToken
func generateToken(u LoginCredentials) (UserToken, error) {

}