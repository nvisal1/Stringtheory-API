package courses

type summary struct {
	id string
	name string
	description string
}

type course struct {
	summary
	lessons []string
}

type dataStore interface {
	getAllCourses() ([]course, error)
}

type httpAdapter interface {
	initializeAdapter()
}