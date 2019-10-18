package service_communication

import . "Stringtheory-API/shared"

type ServiceCommunicator struct {}

func NewServiceCommunicator() *ServiceCommunicator {
	return &ServiceCommunicator{}
}

func (serviceCommunicator ServiceCommunicator) GetCourse(courseID string) (Course, error) {

}

func (serviceCommunicator ServiceCommunicator) GetExercise(exerciseID string) (Exercise, error) {

}

func (serviceCommunicator ServiceCommunicator) GetLesson(lessonID string) (Lesson, error) {

}

func (serviceCommunicator ServiceCommunicator) GetUser(username string) (User, error) {

}