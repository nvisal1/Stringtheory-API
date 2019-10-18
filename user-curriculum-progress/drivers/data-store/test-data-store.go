package data_store

import (
	"Stringtheory-API/user-curriculum-progress"
	"time"
)

type stubMongoDataStore struct {}

func (smds stubMongoDataStore )createCompletedCourse(cc user_curriculum_progress.completedCourse) error {
	return nil
}

func (smds stubMongoDataStore) createCompletedLesson(cL user_curriculum_progress.completedLesson) error {
	return nil
}

func (smds stubMongoDataStore) createCompletedExercise(cE user_curriculum_progress.completedExercise) error {
	return nil
}

func (smds stubMongoDataStore) getCompletedLesson(uN string) (user_curriculum_progress.completedLesson, error) {
	cL := user_curriculum_progress.completedLesson{
		"test_username",
		"test_lesson_id",
		time.Now(),
	}
	return cL, nil
}

func (smds stubMongoDataStore) getCompletedCourse(uN string) (user_curriculum_progress.completedCourse, error) {
	cL := user_curriculum_progress.completedCourse{
		"test_username",
		"test_lesson_id",
		time.Now(),
	}
	return cL, nil
}