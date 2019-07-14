package user_authentication

type User struct {
	username string
	name string
	email string
}

type LoginCredentials struct {
	username string
	password string
}

type UserToken struct {
	token string
}