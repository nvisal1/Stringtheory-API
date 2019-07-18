package user_authentication

import "golang.org/x/crypto/bcrypt"

// comparePasswords is responsible for comparing
// the user's plaintext password to
// the hashed password
func comparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}

func encryptPassword(p string) (string, error) {
	byteP := []byte(p)
	h, err := bcrypt.GenerateFromPassword(byteP, 10)
	if err != nil {
		return "", err
	}
	return string(h), nil
}