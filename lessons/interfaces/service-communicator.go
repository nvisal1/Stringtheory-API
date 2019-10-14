package interfaces

import . "Stringtheory-API/shared"

type ServiceCommunicator interface {
	GetCourse(courseID string) (Course, error)
}
