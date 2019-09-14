package shared

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"strings"
)

 func Authenticate (h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			setupResponse(&w, r)
			if r.Method == "OPTIONS" {
				return
			}
			key, exists := os.LookupEnv("KEY")
			if exists {
				head := r.Header.Get("Authorization")
				if len(strings.Split(head, " ")) != 2 {
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
					h(w, r.WithContext(ctx), p)
					return
				} else {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError) + " Secret not set", http.StatusInternalServerError)
				return
			}
		}
 }

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
