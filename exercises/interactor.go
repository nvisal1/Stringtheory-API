package exercises

import "stringtheory/shared"

func loadLessonExercises(lI string) ([]shared.Exercise, error) {
	es, err := sm.DataStore().getLessonExercises(lI)
	if err != nil {
		return es, err
	}
	return es, err
}

func loadExercise(eI string) (shared.Exercise, error) {
	e, err := sm.DataStore().getExercise(eI)
	return e, err
}
