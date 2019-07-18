package user_authentication

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"stringtheory/shared"
	"time"
)

func generateToken(u shared.SecureUser) (string, error) {
	key, exists := os.LookupEnv("KEY")
	if exists {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = jwt.MapClaims{
			"exp":  time.Now().Add(time.Hour * 72).Unix(),
			"iat":  time.Now().Unix(),
			"user": u,
		}
		byteKey := []byte(key)
		tokenString, err := token.SignedString(byteKey)
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
	return "", errors.New("Key is not set correctly")
}
