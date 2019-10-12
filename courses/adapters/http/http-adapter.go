package http

import (
	. "Stringtheory-API/courses/usecases/load-all-courses"
	"Stringtheory-API/drivers/router"
	"Stringtheory-API/shared"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ModuleHttpAdapter struct {}

// InitializeAdapter is required by the shared/Adapters
// The function is responsible for setting http routes
// for this module
func (mha ModuleHttpAdapter) InitializeAdapter() {
	r := router.GetRouter()
	r.GET("/courses", shared.Authenticate(mha.getCourses))
}

// getCourses handles incoming requests to the /courses route.
// The function ensures that the method used in the request
// is GET. If it is not, the function throws an error.
// If it is, the function will call loadAllCourses in the module
// controller. The response from loadAllCourses is converted to a JSON
// object and returned to the client.
func (mha ModuleHttpAdapter) getCourses(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c, err := LoadAllCourses()
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError) + " " + err.Error(), http.StatusInternalServerError)
		return
	}
	e, err := json.Marshal(c)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError) + " " + err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(e)
}

