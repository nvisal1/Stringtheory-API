package exercises

import (
	database2 "Stringtheory-API/drivers/database"
	"errors"
	"log"
	"os"
)

var sm serviceModule

// InitializeModule is an exported function.
//
// This function looks for SERVICE_ENVIRONMENT
// environment variable. It uses this variable to
// determine whether or not to load test stubs or
// production code.
func InitializeModule() {
	ds, err := setDataStore()
	if err != nil {
		log.Fatal(err)
	}
	sm = serviceModule{
		moduleHttpAdapter{},
		exercisesAdapter{},
		ds,
	}
	sm.ha.InitializeAdapter()
	sm.ia.InitializeAdapter()
}

func setDataStore() (dataStore, error) {
	se, exists := os.LookupEnv("SERVICE_ENVIRONMENT")
	if exists {
		var dataStore dataStore
		switch se {
		case "production":
		case "development":
			dataStore = moduleMongoDataStore{
				database2.GetConnection().Db,
			}
			break
		case "test":
			dataStore = stubMongoDataStore{}
			break
		default:
			log.Fatal("Service environment property is not set correctly")
		}
		return dataStore, nil
	}
	return stubMongoDataStore{}, errors.New("SERVICE_ENVIRONMENT not set")
}
