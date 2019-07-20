package lessons

type lesson struct {
	id string
	name string
	exercises []string
	description string
}

type dataStore interface {
	getCourseLessons(cL []string) ([]lesson, error)
}

type httpAdapter interface {
	initializeAdapter()
}

