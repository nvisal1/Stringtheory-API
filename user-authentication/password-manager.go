package user_authentication

import "golang.org/x/crypto/bcrypt"

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

func encryptPassword(p string) string {

}