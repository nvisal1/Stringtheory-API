package drivers

import (
	"errors"
	"Stringtheory-API/shared"
)

type StubMongoDataStore struct {}

func (smds StubMongoDataStore) GetAllCourses() ([]shared.Course, error) {
	courses := make([]shared.Course, 0)
	c := shared.Course{
		"test",
		"test",
		"test",
		"test",
	}
	res := append(courses, c)
	return res, nil
}

func (smds StubMongoDataStore) GetCourse(cI string) (shared.Course, error) {
	c := shared.Course{
		"test",
		"test",
		"test",
		"test",
	}

	if cI != "" {
		return c, nil
	}
	return c, errors.New("Not found")
}