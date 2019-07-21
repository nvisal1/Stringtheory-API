package courses

import "stringtheory/shared"

type dataStore interface {
	getAllCourses() ([]shared.Course, error)
	getCourse(cI string) (shared.Course, error)
}

type httpAdapter interface {
	initializeAdapter()
}