package user_management

import "stringtheory/shared"

type stubMongoDataStore struct {}

func (mmds stubMongoDataStore) getUser(un string) (shared.User, error) {
	su := shared.SecureUser{
		Username: "test",
		Email: "test",
		Name: "test",
	}
	u := shared.User{
		SecureUser: su,
		Password: "test",
	}
	return u, nil
}

func (mmds stubMongoDataStore) createUser(u shared.User) error {
	return nil
}
