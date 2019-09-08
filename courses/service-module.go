package courses

import (
	"Stringtheory-API/shared"
)

type serviceModule struct {
	ha shared.Adapter
	ia shared.Adapter
	ds dataStore
}