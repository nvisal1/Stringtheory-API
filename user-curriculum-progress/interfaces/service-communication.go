package interfaces

import (
	"Stringtheory-API/shared"
)

type ServiceCommunicator interface {
	GetCourse(courseID string) (*shared.Course, error)
	GetExercise(exerciseID string) (*shared.Exercise, error)
	GetLesson(lessonID string) (*shared.Lesson, error)
	GetUser(username string) (*shared.User, error)
}