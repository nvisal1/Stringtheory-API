package courses

import "Stringtheory-API/shared"

type dataStore interface {
	getAllCourses() ([]shared.Course, error)
	getCourse(cI string) (shared.Course, error)
}