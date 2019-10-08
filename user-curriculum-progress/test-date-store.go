package user_curriculum_progress

import "time"

type stubMongoDataStore struct {}

func (smds stubMongoDataStore )createCompletedCourse(cc completedCourse) error {
	return nil
}

func (smds stubMongoDataStore) createCompletedLesson(cL completedLesson) error {
	return nil
}

func (smds stubMongoDataStore) createCompletedExercise(cE completedExercise) error {
	return nil
}

func (smds stubMongoDataStore) getCompletedLesson(uN string) (completedLesson, error) {
	cL := completedLesson{
		"test_username",
		"test_lesson_id",
		time.Now(),
	}
	return cL, nil
}

func (smds stubMongoDataStore) getCompletedCourse(uN string) (completedCourse, error) {
	cL := completedCourse{
		"test_username",
		"test_lesson_id",
		time.Now(),
	}
	return cL, nil
}