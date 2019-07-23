package courses

import (
	"log"
	"stringtheory/shared"
)

type serviceModule struct {
	ha shared.Adapter
	ia shared.Adapter
	ds dataStore
	tds dataStore
	se string
}

func (sm serviceModule) DataStore() dataStore {
	var dataStore dataStore
	switch sm.se {
	case "production":
	case "development":
		dataStore = sm.ds
		break
	case "test":
		dataStore = sm.tds
		break
	default:
		log.Fatal("Service environment property is not set correctly")
	}
	return dataStore
}