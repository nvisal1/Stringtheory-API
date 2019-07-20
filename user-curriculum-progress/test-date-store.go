package user_curriculum_progress

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