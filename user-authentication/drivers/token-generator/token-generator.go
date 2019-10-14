package token_generator

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"Stringtheory-API/shared"
	"time"
)

type JWTGenerator struct {}

func NewJWTGenerator() *JWTGenerator {
	return &JWTGenerator{}
}

func (jwtGenerator JWTGenerator) GenerateToken(secureUser shared.SecureUser) (string, error) {
	key, exists := os.LookupEnv("KEY")
	if exists {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = jwt.MapClaims{
			"exp":  time.Now().Add(time.Hour * 72).Unix(),
			"iat":  time.Now().Unix(),
			"user": secureUser,
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
