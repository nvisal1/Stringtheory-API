package interfaces

import (
	. "Stringtheory-API/shared"
)

type ServiceCommunicator interface {
	GetCourse(courseID string) (Course, error)
	GetExercise(exerciseID string) (Exercise, error)
	GetLesson(lessonID string) (Lesson, error)
	GetUser(username string) (User, error)
}