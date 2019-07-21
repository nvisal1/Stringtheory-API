package shared

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"os"
	"strings"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Authenticate() Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key, exists := os.LookupEnv("KEY")
			if exists {
				head := r.Header.Get("Authorization")
				t := strings.Split(head, " ")[1]
				jwtSecret := []byte(key)
				if t == "" {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
				token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
						return nil, nil
					}
					return jwtSecret, nil
				})
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					var decodedToken SecureUser
					mapstructure.Decode(claims, &decodedToken)
					ctx := context.WithValue(r.Context(), "User", decodedToken)
					h.ServeHTTP(w, r.WithContext(ctx))
					return
				} else {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError) + " Secret not set", http.StatusInternalServerError)
				return
			}
		})
	}
}
