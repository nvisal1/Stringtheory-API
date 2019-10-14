package interfaces

type PasswordService interface {
	ComparePasswords(hashedPassword string, plainTextPassword string) bool
	EncryptPassword(plainTextPassword string) (string, error)
}
