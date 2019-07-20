package exercises

import "log"

type serviceModule struct {
	ha httpAdapter
	ds dataStore
	tds dataStore
	se string
}

func (sm serviceModule) DataStore() dataStore {
	var dataStore dataStore
	switch sm.se {
	case "production":
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