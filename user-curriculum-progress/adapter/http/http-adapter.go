package http

import (
	"Stringtheory-API/drivers/router"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"Stringtheory-API/shared"
)

type ModuleHttpAdapter struct {}

func NewHttpAdapter() *ModuleHttpAdapter {
	return &ModuleHttpAdapter{}
}

func (adapter ModuleHttpAdapter) InitializeAdapter() {
	r := router.GetRouter()
	r.POST(
		"/users/:userId/courses/:courseId/lessons/:lessonId/exercises/:exerciseId/completion",
		shared.Authenticate(adapter.completeExercise))
}

func (adapter ModuleHttpAdapter) completeExercise(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}