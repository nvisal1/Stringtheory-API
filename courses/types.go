package courses

type dataStore interface {
	getAllCourses() ([]course, error)
	getCourse(cI string) (course, error)
}

type httpAdapter interface {
	initializeAdapter()
}