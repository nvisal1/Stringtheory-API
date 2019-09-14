package user_curriculum_progress

import (
	"Stringtheory-API/drivers/router"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"Stringtheory-API/shared"
)

type moduleHttpAdapter struct {}

func (mha moduleHttpAdapter) InitializeAdapter() {
	r := router.GetRouter()
	r.POST(
		"/users/:userId/courses/:courseId/lessons/:lessonId/exercises/:exerciseId/complete",
		shared.Authenticate(mha.completeExercise))
}

func (mha moduleHttpAdapter) completeExercise(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}