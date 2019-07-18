package user_authentication

import (
	"stringtheory/shared"
	user_management "stringtheory/user-management"
)


func getUser(un string) (shared.User, error) {
	u, err := user_management.GetInstance().GetUser(un)
	return u, err
}