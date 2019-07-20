package exercises

func loadLessonExercises(lI string) ([]exercise, error) {
	es, err := sm.DataStore().getLessonExercises(lI)
	if err != nil {
		return es, err
	}
	return es, err
}

func loadExercise(eI string) (exercise, error) {
	e, err := sm.DataStore().getExercise(eI)
	return e, err
}
