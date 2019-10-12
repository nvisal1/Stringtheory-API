package interfaces

import "Stringtheory-API/shared"

type Datastore interface {
	GetAllCourses() ([]shared.Course, error)
	GetCourse(courseID string) (shared.Course, error)
}