package user_curriculum_progress

import "time"

type completedCourse struct {
	userId string
	courseId string
	date time.Time
}

type completedLesson struct {
	userId string
	lessonId string
	date time.Time
}

type completedExercise struct {
	userId string
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
