package courses

import (
	"Stringtheory-API/drivers/router"
	"Stringtheory-API/shared"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type moduleHttpAdapter struct {}

// InitializeAdapter is required by the shared/Adapters
// The function is responsible for setting http routes
// for this module
func (mha moduleHttpAdapter) InitializeAdapter() {
	r := router.GetRouter()
	r.GET("/courses", shared.Authenticate(mha.getCourses))
}

// getCourses handles incoming requests to the /courses route.
// The function ensures that the method used in the request
// is GET. If it is not, the function throws an error.
// If it is, the function will call loadAllCourses in the module
// controller. The response from loadAllCourses is converted to a JSON
// object and returned to the client.
func (mha moduleHttpAdapter) getCourses(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c, err := loadAllCourses()
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

// getCourse handles incoming requests to /courses/:courseId.
// This functionality is not implemented.
func (mha moduleHttpAdapter) getCourse(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound) + " Hint: try making a GET request to this endpoint", http.StatusNotFound)
		return
	}
}

