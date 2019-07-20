package lessons

func loadCourseLessons(cI string) ([]lesson, error) {
	// get course
	c := getCourse(cI)
	l, err := sm.DataStore().getCourseLessons(c.lessons)
	if err != nil {
		return l, err
	}
	return l, nil
}
