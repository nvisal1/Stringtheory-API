package interfaces

import . "Stringtheory-API/user-curriculum-progress/types"

type Datastore interface {
	CreateCompletedCourse(completedCourse *CompletedCourse) error
	CreateCompletedLesson(completedLesson CompletedLesson) error
	CreateCompletedExercise(completedExercise CompletedExercise) error
	GetCompletedLesson(username string) (CompletedLesson, error)
	GetCompletedCourse(username string) (CompletedCourse, error)
}
