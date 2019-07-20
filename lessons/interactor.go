package lessons

func loadCourseLessons(cI string) ([]lesson, error) {
	// get course to check existence
	c := getCourse(cI)
	l, err := sm.DataStore().getCourseLessons(cI)
	if err != nil {
		return l, err
	}
	return l, nil
}

func loadLesson(lI string) (lesson, error) {
	l, err := sm.DataStore().getLesson(lI)
	return l, err
}
