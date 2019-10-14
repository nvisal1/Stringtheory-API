package register

import (
	"Stringtheory-API/shared"
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

func TestRegister(t *testing.T) {
	Convey("When Register is called", t, func() {
		stubUser := shared.User{
			Username: "test_username",
			Email: "test_email",
			Name: "test_name",
			Password: "test_password",
		}

		userToken, err := Register(stubUser)

		Convey("and valid credentials are provided", func() {
			expectedResponse := types.UserToken{Token: "test_token_strings"}

			Convey("a UserToken should be returned without error", func() {
				So(userToken, ShouldResemble, expectedResponse)
				So(err, ShouldBeNil)
			})
		})
	})
}

