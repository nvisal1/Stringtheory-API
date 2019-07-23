package courses

import "stringtheory/shared"

type stubMongoDataStore struct {}

func (smds stubMongoDataStore) getAllCourses() ([]shared.Course, error) {
	courses := make([]shared.Course, 1)
	c := shared.Course{
		"test",
		"test",
		"test",
	}
	res := append(courses, c)
	return res, nil
}

func (smds stubMongoDataStore) getCourse(cI string) (shared.Course, error) {
	c := shared.Course{
		"test",
		"test",
		"test",
	}
	return c, nil
}