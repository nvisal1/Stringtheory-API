package http

import (
	"Stringtheory-API/drivers/router"
	. "Stringtheory-API/lessons/usecases/load-course-lessons"
	"Stringtheory-API/shared"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ModuleHttpAdapter struct {}

func (adapter ModuleHttpAdapter) InitializeAdapter() {
	r := router.GetRouter()
	r.GET("/courses/:courseId/lessons", shared.Authenticate(adapter.getCourseLessons))
}

func (adapter ModuleHttpAdapter) getCourseLessons(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	courseID := p.ByName("courseId")
	c, err := LoadCourseLessons(courseID)
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

