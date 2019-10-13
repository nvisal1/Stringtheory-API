package load_lesson_exercises

import (
	. "Stringtheory-API/exercises/service-module"
	"Stringtheory-API/shared"
)

// loadAllLessonExercises is responsible for asking
// the database for all exercises that belong to a given lesson.
// This function returns a slice of exercises
// and an error
func LoadLessonExercises(lessonID string) ([]shared.Exercise, error) {
	// TODO: Check if specified lesson exists
	lessonExercises, err := ExercisesModule.Datastore.QueryLessonExercises(lessonID)
	return lessonExercises, err
}
