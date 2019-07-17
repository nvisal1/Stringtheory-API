package user_authentication

type loginCredentials struct {
	username string
	password string
}

type newUser struct {
	username string
	name string
	email string
	password string
}

type userToken struct {
	token string
}

type userManagementContract struct {
	username string
	name string
	email string
}

type httpAdapter interface {
	initializeAdapter()
}




