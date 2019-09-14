package user_management

import (
	"Stringtheory-API/drivers/router"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"Stringtheory-API/shared"
)

type moduleHttpAdapter struct {}

func (mha moduleHttpAdapter) InitializeAdapter() {
	r := router.GetRouter()
	r.PATCH("/users", shared.Authenticate(mha.edit))
	r.DELETE("/users", shared.Authenticate((mha.delete)))
}

func (mha moduleHttpAdapter) edit(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}

func (mha moduleHttpAdapter) delete(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}