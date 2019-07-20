package user_curriculum_progress

import "time"

func completeCourse(uI string, cI string) error {
	// get course to verify existence

	// get user to verify existence

	// make sure that the user hasn't completed the course already

	cC := completedCourse{
		uI,
		cI,
		time.Now(),
	}
	err := sm.DataStore().createCompletedCourse(cC)
	return err
}

func completeLesson(uI string, lI string) error {
	// get lesson to verify existence

	// get user to verify existence

	// make sure that the user hasn't completed the lesson already

	cL := completedLesson{
		uI,
		lI,
		time.Now(),
	}

	err := sm.DataStore().createCompletedLesson(cL)
	return err
}

func completeExercise(uI string, eI string, score float64) error {
	// get exercise to verify existence

	// get user to verify existence

	cE := completedExercise{
		uI,
		eI,
		time.Now(),
		score,
	}

	err := sm.DataStore().createCompletedExercise(cE)
	return err
}