package service_communication

import (
	. "Stringtheory-API/courses/usecases/load-course"
	"Stringtheory-API/shared"
)

type ServiceCommunicator struct {}

func (serviceCommunicator ServiceCommunicator)GetCourse(courseID string) (shared.Course, error) {
	c, err := LoadCourse(courseID)
	return c, err
}