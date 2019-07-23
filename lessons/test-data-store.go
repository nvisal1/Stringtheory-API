package lessons

import "stringtheory/shared"

type stubMongoDataStore struct {}

func (smds stubMongoDataStore) getCourseLessons(cI string) ([]shared.Lesson, error) {
	es := make([]shared.Lesson, 3)
	e := shared.Lesson {
		"test",
		"test",
		1,
		"test",
		"test",
		true,
	}
	return append(es, e), nil
}

func (smds stubMongoDataStore) getLesson(lI string) (shared.Lesson, error) {
	l := shared.Lesson {
		"test",
		"test",
		1,
		"test",
		"test",
		true,
	}

	return l, nil
}