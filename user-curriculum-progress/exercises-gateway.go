package user_curriculum_progress

import (
	"Stringtheory-API/exercises"
	"Stringtheory-API/shared"
)

func getExercise(eI string) (shared.Exercise, error) {
	e, err := exercises.GetInstance().LoadExercise(eI)
	return e, err
}