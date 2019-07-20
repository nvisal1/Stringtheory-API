package courses

import (
	"encoding/json"
	"net/http"
)

type moduleHttpAdapter struct {}

func (mha moduleHttpAdapter) initializeAdapter() {
	http.HandleFunc("/courses", mha.getCourses)
	http.HandleFunc("/courses/:courseId", mha.getCourse)
}

func (mha moduleHttpAdapter) getCourses(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		c, err := loadAllCourses()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError) + " " + err.Error(), http.StatusInternalServerError)
			return
		}
		e, err := json.Marshal(c)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError) + " " + err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(e)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound) + " Hint: try making a GET request to this endpoint", http.StatusNotFound)
		return
	}
}


func (mha moduleHttpAdapter) getCourse(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound) + " Hint: try making a GET request to this endpoint", http.StatusNotFound)
		return
	}
}

