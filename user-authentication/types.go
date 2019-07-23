package user_authentication

type loginCredentials struct {
	Username string
	Password string
}

type userToken struct {
	Token string
}

type httpAdapter interface {
	initializeAdapter()
}




