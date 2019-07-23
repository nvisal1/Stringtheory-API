package lessons

import (
	"encoding/json"
	"net/http"
	"stringtheory/shared"
)

type moduleHttpAdapter struct {}

func (mha moduleHttpAdapter) InitializeAdapter() {
	http.HandleFunc("/courses/:courseId/lessons", shared.Authenticate(mha.getCourseLessons))
}

func (mha moduleHttpAdapter) getCourseLessons(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		cI := req.FormValue("courseId")
		c, err := loadCourseLessons(cI)
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
