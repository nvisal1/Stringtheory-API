package password_service

type StubPasswordService struct {}

func NewStubPasswordService() *StubPasswordService {
	return &StubPasswordService{}
}

func (stubPasswordService StubPasswordService) ComparePasswords(hashedPassword string, plainTextPassword string) bool {
	return true
}

func (stubPasswordService StubPasswordService) EncryptPassword(plainTextPassword string) (string, error) {
	return "test_encrypted_password", nil
}