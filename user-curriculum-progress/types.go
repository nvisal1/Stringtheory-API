package user_curriculum_progress

import "time"

type completedCourse struct {
	username string
	courseId string
	date time.Time
}

type completedLesson struct {
	username string
	lessonId string
	date time.Time
}

type completedExercise struct {
	username string
	exerciseId string
	date time.Time
	score float64
}


type dataStore interface {
	createCompletedCourse(cc completedCourse) error
	createCompletedLesson(cL completedLesson) error
	createCompletedExercise(cE completedExercise) error
}

type httpAdapter interface {
	initializeAdapter()
}
