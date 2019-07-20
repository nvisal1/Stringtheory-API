package courses

type course struct {
	id string
	name string
	description string
}

type dataStore interface {
	getAllCourses() ([]course, error)
	getCourse(cI string) (course, error)
}

type httpAdapter interface {
	initializeAdapter()
}