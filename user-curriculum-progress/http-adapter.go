package user_curriculum_progress

import "net/http"

type moduleHttpAdapter struct {}

func (mha moduleHttpAdapter) InitializeAdapter() {
	http.HandleFunc("/users/:userId/courses/:courseId/lessons/:lessonId/exercises/:exerciseId/complete", mha.completeExercise)
}

func (mha moduleHttpAdapter) completeExercise(w http.ResponseWriter, req *http.Request) {

}