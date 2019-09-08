package user_curriculum_progress

import (
	"Stringtheory-API/shared"
	user_management "Stringtheory-API/user-management"
)

func getUser(un string) (shared.User, error) {
	u, err := user_management.GetInstance().GetUser(un)
	return u, err
}