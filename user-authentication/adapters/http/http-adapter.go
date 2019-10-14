package http

import (
	"Stringtheory-API/drivers/router"
	"Stringtheory-API/shared"
	"Stringtheory-API/user-authentication/types"
	. "Stringtheory-API/user-authentication/usecases/login"
	. "Stringtheory-API/user-authentication/usecases/register"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type HttpAdapter struct{}

func NewHttpAdapter() *HttpAdapter {
	return &HttpAdapter{}
}

func (adapter HttpAdapter) InitializeAdapter() {
	r := router.GetRouter()
	r.POST("/login", adapter.login)
	r.POST("/register", adapter.register)
}

// login checks to ensure that the incoming request is a
// POST request. If it is not a POST request, the function
// will return a 404 error.
//
// If the request is a POST request, the function
// will read the request body. The request body
// should contain a JSON object with the username
// and password properties. This JSON object is
// decoded by the json package and stored in a
// variable with the type LoginCredentials.
// This variable is then passed to the generateToken
// function.
//
// If no errors are returned from the generateToken function,
// the variable of type UserToken is encoded and returned
// to the client.
func (adapter HttpAdapter) login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}

	var loginCredentials types.LoginCredentials
	err = json.Unmarshal(body, &loginCredentials)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}

	userToken, err := Login(loginCredentials)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized)+" "+err.Error(), http.StatusUnauthorized)
		return
	}

	response, err := json.Marshal(userToken)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(response)
}

func (adapter HttpAdapter) register(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user shared.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}

	userToken, err := Register(user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest)+" "+err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(userToken)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(response)
}