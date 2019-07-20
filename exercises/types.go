package exercises

type exercise struct {
	id string
	name string
	order int
	notes []string
	description string
	lessonId string
	hasNext bool
}

type dataStore interface {
	getLessonExercises(lI string) ([]exercise, error)
	getExercise(eI string) (exercise, error)
}

type httpAdapter interface {
	initializeAdapter()
}