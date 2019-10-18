package complete_lesson

import (
	user_curriculum_progress "Stringtheory-API/user-curriculum-progress"
	"time"
)

func CompleteLesson(uN string, lI string) error {
	// get lesson to verify existence
	l, err := user_curriculum_progress.getLesson(lI)
	if err != nil {
		return err
	}
	// make sure that the user hasn't completed the lesson already

	cL := user_curriculum_progress.completedLesson{
		uN,
		lI,
		time.Now(),
	}

	err = user_curriculum_progress.sm.DataStore().createCompletedLesson(cL)
	if err != nil {
		return err
	}

	// check if all lessons were completed
	// if yes - complete course
	if !l.HasNext {
		err = completeCourse(uN, l.CourseId)
		if err != nil {
			return err
		}
	}
	return nil
}