package exercises

import (
	"errors"
	"stringtheory/shared"
)

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

	if lI != "" {
		return append(es, e), nil
	}
	return append(es, e), errors.New("Not found")
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
	if eI != "" {
		return e, nil
	}
	return e, errors.New("Not found")
}