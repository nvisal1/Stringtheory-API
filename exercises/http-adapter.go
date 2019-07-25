package exercises

import (
	"encoding/json"
	"net/http"
	"Stringtheory-API/shared"
)

type moduleHttpAdapter struct {}

func (mha moduleHttpAdapter) InitializeAdapter() {
	http.HandleFunc("courses/:courseId/lessons/:lessonId/exercises", shared.Authenticate(mha.getLessonExercises))
}

func (mha moduleHttpAdapter) getLessonExercises(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		lI := req.FormValue("lessonId")
		c, err := loadLessonExercises(lI)
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
