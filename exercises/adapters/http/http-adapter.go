package http

import (
	"Stringtheory-API/drivers/router"
	. "Stringtheory-API/exercises/usecases/load-lesson-exercises"
	"Stringtheory-API/shared"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ModuleHttpAdapter struct {}

func (adapter ModuleHttpAdapter) InitializeAdapter() {
	r := router.GetRouter()
	r.GET("/courses/:courseId/lessons/:lessonId/exercises", shared.Authenticate(adapter.getLessonExercises))
}

func (mha ModuleHttpAdapter) getLessonExercises(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	lessonID := p.ByName("lessonId")
	exercises, err := LoadLessonExercises(lessonID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError) + " " + err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(exercises)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError) + " " + err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(response)
}
