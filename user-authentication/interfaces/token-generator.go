package interfaces

import "Stringtheory-API/shared"

type TokenGenerator interface {
	GenerateToken(secureUser shared.SecureUser) (string, error)
}
