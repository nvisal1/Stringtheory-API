package types

import "time"

type CompletedCourse struct {
	Username string
	CourseId string
	Date time.Time
}

type CompletedLesson struct {
	Username string
	LessonId string
	Date time.Time
}

type CompletedExercise struct {
	Username string
	ExerciseId string
	Date time.Time
	Score float64
}

func NewCompletedCourse(username string, courseID string, date time.Time) *CompletedCourse {
	return &CompletedCourse{
		Username: username,
		CourseId: courseID,
		Date: date,
	}
}

func NewCompletedLesson(username string, lessonID string, date time.Time) *CompletedLesson {
	return &CompletedLesson{
		Username: username,
		LessonId: lessonID,
		Date:     date,
	}
}

func NewCompletedExercise(username string, exerciseID string, date time.Time, score float64) *CompletedExercise {
	return &CompletedExercise{
		Username:   username,
		ExerciseId: exerciseID,
		Date:       date,
		Score:      score,
	}
}

