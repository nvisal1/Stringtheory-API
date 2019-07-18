package user_authentication

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




