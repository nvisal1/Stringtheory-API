package lessons

type stubMongoDataStore struct {}

func (smds stubMongoDataStore) getCourseLessons(cI string) ([]lesson, error) {
	es := make([]lesson, 3)
	e := lesson {
		"test",
		"test",
		1,
		"test",
		"test",
		true,
	}
	return append(es, e), nil
}

func (smds stubMongoDataStore) getLesson(lI string) (lesson, error) {
	l := lesson {
		"test",
		"test",
		1,
		"test",
		"test",
		true,
	}

	return l, nil
}