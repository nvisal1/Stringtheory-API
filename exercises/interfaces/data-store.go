package interfaces

import "Stringtheory-API/shared"

type Datastore interface {
	QueryLessonExercises(lI string) ([]shared.Exercise, error)
	QueryExercise(eI string) (shared.Exercise, error)
}