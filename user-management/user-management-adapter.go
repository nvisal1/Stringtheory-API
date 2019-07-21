package user_management

import (
	"log"
	"stringtheory/shared"
	"sync"
)

type userManagementAdapter struct {}

var instance userManagementAdapter
var once sync.Once

func (uma userManagementAdapter) InitializeAdapter() {
	once.Do(func() {
		instance = userManagementAdapter{}
	})
}

func (uma userManagementAdapter) GetUser(un string) (shared.User, error) {
	u, err := getUser(un)
	return u, err
}

func (uma userManagementAdapter) CreateUser(u shared.User) error {
	err := createUser(u)
	return err
}

func GetInstance() userManagementAdapter {
	if instance.GetUser == nil {
		log.Fatal("User Management Adapter has not been created")
	}
	return instance
}

