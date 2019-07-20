package courses

func loadAllCourses() ([]course, error) {
	c, err := sm.DataStore().getAllCourses()
	if err != nil {
		return c, err
	}
	return c, nil
}

