package login

import (
	. "Stringtheory-API/user-authentication/drivers/password-service"
	. "Stringtheory-API/user-authentication/drivers/service-communication"
	. "Stringtheory-API/user-authentication/drivers/token-generator"
	. "Stringtheory-API/user-authentication/service-module"
	"Stringtheory-API/user-authentication/types"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("SERVICE_ENVIRONMENT", "test")
	NewUserAuthenticationModule(nil, NewStubPasswordService(), NewStubTokenGenerator(), NewStubServiceCommunicator())
}

func TestLogin(t *testing.T) {
	Convey("When Login is called", t, func() {
		stubLoginCredentials := types.LoginCredentials{Username: "test_username" , Password: "test_password"}
		userToken, err := Login(stubLoginCredentials)

		Convey("and valid credentials are provided", func() {
			expectedResponse := types.UserToken{Token: "test_token_string"}

			Convey("a UserToken should be returned without error", func() {
				So(userToken, ShouldResemble, expectedResponse)
				So(err, ShouldBeNil)
			})
		})
	})
}