package complete_exercise

import (
	user_curriculum_progress "Stringtheory-API/user-curriculum-progress"
	"time"
)

func CompleteExercise(uN string, eI string, score float64) error {
	// get exercise to verify existence
	e, err := user_curriculum_progress.getExercise(eI)
	if err != nil {
		return err
	}
	// get user to verify existence
	_, err = user_curriculum_progress.getUser(uN)
	if err != nil {
		return err
	}

	cE := user_curriculum_progress.completedExercise{
		uN,
		eI,
		time.Now(),
		score,
	}

	err = user_curriculum_progress.sm.DataStore().createCompletedExercise(cE)
	if err != nil {
		return err
	}

	// Check if all exercises were completed
	// if yes - complete lesson
	if !e.HasNext {
		err = completeLesson(uN, e.LessonId)
		if err != nil {
			return err
		}
	}
	return nil
}