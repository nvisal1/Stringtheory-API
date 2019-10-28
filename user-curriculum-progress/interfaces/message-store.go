package interfaces

import . "Stringtheory-API/user-curriculum-progress/types"

type MessageStore interface {
	SendMessage(message *Message) error
	ReceiveMessage() (*Message, error)
}
