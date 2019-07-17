package user_authentication

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type moduleHttpAdapter struct{}

func (mha moduleHttpAdapter) initializeAdapter() {
	http.HandleFunc("/login", mha.login)
	http.HandleFunc("/register", mha.register)
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
func (mha moduleHttpAdapter) login(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		b, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError) + " " + err.Error(), http.StatusInternalServerError)
			return
		}

		var u loginCredentials
		err = json.Unmarshal(b, &u)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError) + " " + err.Error(), http.StatusInternalServerError)
			return
		}

		ut, err := generateToken(u)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized) + " " + err.Error(), http.StatusUnauthorized)
			return
		}

		e, err := json.Marshal(ut)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError) + " " + err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(e)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound) + " Hint: try making a POST request to this endpoint", http.StatusNotFound)
		return
	}
}

func (mha moduleHttpAdapter) register(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}

	var u newUser
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}
}
