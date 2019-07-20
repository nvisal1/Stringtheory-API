package lessons

type lesson struct {
	id string
	name string
	order int
	description string
	courseId string
	hasNext bool
}

type dataStore interface {
	getCourseLessons(cI string) ([]lesson, error)
	getLesson(lI string) (lesson, error)
}

type httpAdapter interface {
	initializeAdapter()
}

