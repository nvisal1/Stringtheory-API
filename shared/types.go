package shared

type SecureUser struct {
	Username string
	Name string
	Email string
}

type User struct {
	SecureUser
	Password string
}


