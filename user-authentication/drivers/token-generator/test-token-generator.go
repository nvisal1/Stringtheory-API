package token_generator

import "Stringtheory-API/shared"

type StubTokenGenerator struct {}

func NewStubTokenGenerator() *StubTokenGenerator {
	return &StubTokenGenerator{}
}

func (stubTokenGenerator StubTokenGenerator) GenerateToken(secureUser shared.SecureUser) (string, error) {
	return "test_token_string", nil
}