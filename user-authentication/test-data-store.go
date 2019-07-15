package user_authentication

type moduleTestDataStore struct{}

func (mtds moduleTestDataStore) getUser(un string) user {
	return user{
		username: "test",
		name: "test",
		email: "test",
	}
}
