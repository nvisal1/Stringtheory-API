package user_curriculum_progress

import "time"

type completedCourse struct {
	Username string
	CourseId string
	Date time.Time
}

type completedLesson struct {
	Username string
	LessonId string
	Date time.Time
}

type completedExercise struct {
	Username string
	ExerciseId string
	Date time.Time
	Score float64
}


type dataStore interface {
	createCompletedCourse(cc completedCourse) error
	createCompletedLesson(cL completedLesson) error
	createCompletedExercise(cE completedExercise) error
	getCompletedLesson(uN string) (completedLesson, error)
	getCompletedCourse(uN string) (completedCourse, error)
}

type httpAdapter interface {
	initializeAdapter()
}
