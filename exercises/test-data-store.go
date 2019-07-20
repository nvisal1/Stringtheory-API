package exercises

type stubMongoDataStore struct {}

func (smds stubMongoDataStore) getLessonExercises(lE []string) ([]exercise, error) {
	es := make([]exercise, 3)
	e := exercise {
		"test",
		"test",
	    0,
		make([]string, 3),
		"test",
		"test",
		false,
	}
	return append(es, e), nil
}

func (smds stubMongoDataStore) getExercise(eI string) (exercise, error) {
	e := exercise {
		"test",
		"test",
		0,
		make([]string, 3),
		"test",
		"test",
		false,
	}
	return e, nil
}