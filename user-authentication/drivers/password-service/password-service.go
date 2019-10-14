package password_service

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordSerivce struct {}

func NewBcryptPasswordService() *BcryptPasswordSerivce {
	return &BcryptPasswordSerivce{}
}

// comparePasswords is responsible for comparing
// the user's plaintext password to
// the hashed password
func (BcryptPasswordSerivce BcryptPasswordSerivce) ComparePasswords(hashedPassword string, plainTextPassword string) bool {
	byteHash := []byte(hashedPassword)
	bytePlain := []byte(plainTextPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}

func (BcryptPasswordSerivce BcryptPasswordSerivce) EncryptPassword(plainTextPassword string) (string, error) {
	byteP := []byte(plainTextPassword)
	h, err := bcrypt.GenerateFromPassword(byteP, 10)
	if err != nil {
		return "", err
	}
	return string(h), nil
}