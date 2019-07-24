package courses

import (
	"stringtheory/shared"
)

type serviceModule struct {
	ha shared.Adapter
	ia shared.Adapter
	ds dataStore
}