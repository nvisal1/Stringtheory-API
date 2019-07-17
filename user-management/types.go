package user_management

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

type dataStore interface {
	getUser(un string) user
}

type httpAdapter interface {
	initializeAdapter()
}