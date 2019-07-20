package courses

type stubMongoDataStore struct {}

func (smds stubMongoDataStore) getAllCourses() ([]course, error) {
	courses := make([]course, 1)
	c := course{
		"test",
		"test",
		"test",
	}
	res := append(courses, c)
	return res, nil
}

func (smds stubMongoDataStore) getCourse(cI string) (course, error) {
	c := course{
		"test",
		"test",
		"test",
	}
	return c, nil
}