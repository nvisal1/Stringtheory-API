package exercises

import "stringtheory/shared"

type stubMongoDataStore struct {}

func (smds stubMongoDataStore) getLessonExercises(lI string) ([]shared.Exercise, error) {
	es := make([]shared.Exercise, 3)
	e := shared.Exercise {
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

func (smds stubMongoDataStore) getExercise(eI string) (shared.Exercise, error) {
	e := shared.Exercise {
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