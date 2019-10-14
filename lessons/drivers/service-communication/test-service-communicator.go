package service_communication

import . "Stringtheory-API/shared"

type StubServiceCommunicator struct {}

func (serviceCommunicator StubServiceCommunicator)GetCourse(courseID string) (Course, error) {
	course := Course{
		ID: "test_ID",
		Name: "test_name",
		Description: "test_description",
		LessonsURI: "test_lessonURI",
	}
	return course, nil
}
