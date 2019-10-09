package interfaces

import "Stringtheory-API/shared"

type Datastore interface {
	getAllCourses() ([]shared.Course, error)
}