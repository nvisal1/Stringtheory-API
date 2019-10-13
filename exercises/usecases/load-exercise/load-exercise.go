package load_exercise

import (
	. "Stringtheory-API/exercises/service-module"
	"Stringtheory-API/shared"
)

// loadExercises is responsible for asking the
// database for a single exercise.
// This function accepts an exercise id and
// returns an exercise and an error
func LoadExercise(exerciseID string) (shared.Exercise, error) {
	exercise, err := ExercisesModule.Datastore.QueryExercise(exerciseID)
	return exercise, err
}
