package shared

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

 func Authenticate (h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key, exists := os.LookupEnv("KEY")
			if exists {
				head := r.Header.Get("Authorization")
				if len(head) != 2 {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
				t := strings.Split(head, " ")[1]
				jwtSecret := []byte(key)
				if t== "" {
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
					user := claims["user"]
					ctx := context.WithValue(r.Context(), "User", user)
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