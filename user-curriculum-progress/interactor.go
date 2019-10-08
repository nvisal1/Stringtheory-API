package user_curriculum_progress

import "time"

func completeCourse(uN string, cI string) error {
	// get course to verify existence
    _, err := getCourse(cI)
    if err != nil {
    	return err
	}
	// make sure that the user hasn't completed the course already


	cC := completedCourse{
		uN,
		cI,
		time.Now(),
	}
	err = sm.DataStore().createCompletedCourse(cC)
	return err
}

func completeLesson(uN string, lI string) error {
	// get lesson to verify existence
	l, err := getLesson(lI)
	if err != nil {
		return err
	}
	// make sure that the user hasn't completed the lesson already

	cL := completedLesson{
		uN,
		lI,
		time.Now(),
	}

	err = sm.DataStore().createCompletedLesson(cL)
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

func completeExercise(uN string, eI string, score float64) error {
	// get exercise to verify existence
	e, err := getExercise(eI)
	if err != nil {
		return err
	}
	// get user to verify existence
	_, err = getUser(uN)
	if err != nil {
		return err
	}

	cE := completedExercise{
		uN,
		eI,
		time.Now(),
		score,
	}

	err = sm.DataStore().createCompletedExercise(cE)
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