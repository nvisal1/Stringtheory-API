package exercises

type dataStore interface {
	getLessonExercises(lI string) ([]exercise, error)
	getExercise(eI string) (exercise, error)
}

type httpAdapter interface {
	initializeAdapter()
}