package exercises

func loadLessonExercises(lI string) ([]exercise, error) {
	// Get Lesson
	l := getLesson(lI)
	es, err := sm.DataStore().getLessonExercises(l.exercises)
	if err != nil {
		return es, err
	}
	return es, err
}
