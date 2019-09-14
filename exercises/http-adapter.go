package exercises

import (
	"Stringtheory-API/drivers/router"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"Stringtheory-API/shared"
)

type moduleHttpAdapter struct {}

func (mha moduleHttpAdapter) InitializeAdapter() {
	r := router.GetRouter()
	r.GET("/courses/:courseId/lessons/:lessonId/exercises", shared.Authenticate(mha.getLessonExercises))
}

func (mha moduleHttpAdapter) getLessonExercises(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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
}
