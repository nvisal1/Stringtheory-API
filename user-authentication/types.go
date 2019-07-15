package user_authentication

type user struct {
	username string
	name string
	email string
	password string
}

type secureUser struct {
	username string
	name string
	email string
}

type loginCredentials struct {
	username string
	password string
}

type userToken struct {
	token string
}

type httpAdapter interface {
	initializeAdapter()
}


type dataStore interface {
	getUser(un string) user
}

