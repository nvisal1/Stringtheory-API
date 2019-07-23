package user_management

import (
	"net/http"
	"stringtheory/shared"
)

type moduleHttpAdapter struct {}

func (mha moduleHttpAdapter) InitializeAdapter() {
	http.HandleFunc("/users", shared.Authenticate()(mha.handleUsers))
}

func (mha moduleHttpAdapter) handleUsers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPatch:
		mha.edit(w, req)
		break
	case http.MethodDelete:
		mha.delete(w, req)
		break
	default:
		http.Error(w, http.StatusText(http.StatusNotFound) + " Hint: try making a POST, PATCH or DELETE request to this endpoint", http.StatusNotFound)
		return
	}
}

func (mha moduleHttpAdapter) edit(w http.ResponseWriter, req *http.Request) {

}

func (mha moduleHttpAdapter) delete(w http.ResponseWriter, req *http.Request) {

}