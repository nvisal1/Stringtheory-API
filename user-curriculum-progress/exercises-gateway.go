package user_curriculum_progress

import (
	"stringtheory/exercises"
	"stringtheory/shared"
)

func getExercise(eI string) (shared.Exercise, error) {
	e, err := exercises.GetInstance().LoadExercise(eI)
	return e, err
}