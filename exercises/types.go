package exercises

import "Stringtheory-API/shared"

type dataStore interface {
	getLessonExercises(lI string) ([]shared.Exercise, error)
	getExercise(eI string) (shared.Exercise, error)
}