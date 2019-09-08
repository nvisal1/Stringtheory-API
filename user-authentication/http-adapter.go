package user_authentication

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"Stringtheory-API/shared"
)

type moduleHttpAdapter struct{}

func (mha moduleHttpAdapter) InitializeAdapter() {
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
	setupResponse(&w, req)
	if req.Method == "OPTIONS" {
		return
	}
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

		ut, err := processLogin(u)
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
	setupResponse(&w, req)
	if req.Method == "OPTIONS" {
		return
	}
	if req.Method == http.MethodPost {
		b, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
			return
		}

		var u shared.User
		err = json.Unmarshal(b, &u)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
			return
		}

		ut, err := processRegistration(u)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest)+" "+err.Error(), http.StatusBadRequest)
			return
		}

		e, err := json.Marshal(ut)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.Write(e)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound) + " Hint: try making a POST request to this endpoint", http.StatusNotFound)
		return
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
