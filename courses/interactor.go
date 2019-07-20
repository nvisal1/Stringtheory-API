package courses

func loadAllCourses() ([]course, error) {
	c, err := sm.DataStore().getAllCourses()
	if err != nil {
		return c, err
	}
	return c, nil
}

func loadCourse(cI string) (course, error) {
	c, err := sm.DataStore().getCourse(cI)
	return c, err
}