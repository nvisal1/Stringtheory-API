package exercises

import "Stringtheory-API/shared"

// loadAllLessonExercises is responsible for asking
// the database for all exercises that belong to a given lesson.
// This function returns a slice of exercises
// and an error
func loadLessonExercises(lI string) ([]shared.Exercise, error) {
	es, err := sm.ds.getLessonExercises(lI)
	if err != nil {
		return es, err
	}
	return es, err
}

// loadExercises is responsible for asking the
// database for a single exercise.
// This function accepts an exercise id and
// returns an exercise and an error
func loadExercise(eI string) (shared.Exercise, error) {
	e, err := sm.ds.getExercise(eI)
	return e, err
}
