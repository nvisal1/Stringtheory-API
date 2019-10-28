package service_communication

import (
	courseService "Stringtheory-API/courses/usecases/load-course"
	exerciseService "Stringtheory-API/exercises/usecases/load-exercise"
	lessonService "Stringtheory-API/lessons/usecases/load-lesson"
	. "Stringtheory-API/shared"
	"errors"
)

type ServiceCommunicator struct {}

func NewServiceCommunicator() *ServiceCommunicator {
	return &ServiceCommunicator{}
}

func (serviceCommunicator ServiceCommunicator) GetCourse(courseID string) (*Course, error) {
	course, err := courseService.LoadCourse(courseID)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (serviceCommunicator ServiceCommunicator) GetExercise(exerciseID string) (*Exercise, error) {
	exercise, err := exerciseService.LoadExercise(exerciseID)
	if err != nil {
		return nil, err
	}
	return &exercise, nil
}

func (serviceCommunicator ServiceCommunicator) GetLesson(lessonID string) (*Lesson, error) {
	lesson, err := lessonService.LoadLesson(lessonID)
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (serviceCommunicator ServiceCommunicator) GetUser(username string) (*User, error) {
	return nil, errors.New("not implemented")
}