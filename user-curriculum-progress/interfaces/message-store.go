package interfaces

type MessageStore interface {
	SendMessage() error
	ReceiveMessages() (string, error)
}
