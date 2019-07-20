package exercises

type exercise struct {
	id string
	name string
	notes []string
	description string
}

type dataStore interface {
	getLessonExercises(lE []string) ([]exercise, error)
}

type httpAdapter interface {
	initializeAdapter()
}