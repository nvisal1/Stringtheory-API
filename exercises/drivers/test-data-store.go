package drivers

import (
	"Stringtheory-API/shared"
)

type StubMongoDataStore struct {}

func (datastore StubMongoDataStore) QueryLessonExercises(lessonID string) ([]shared.Exercise, error) {
	exercises := make([]shared.Exercise, 3)
	exercise := shared.Exercise {
		"test",
		"test",
	    0,
		make([]string, 3),
		"test",
		"test",
		false,
	}
	return append(exercises, exercise), nil
}

func (datastore StubMongoDataStore) QueryExercise(exerciseID string) (shared.Exercise, error) {
	exercise := shared.Exercise {
		"test",
		"test",
		0,
		make([]string, 3),
		"test",
		"test",
		false,
	}
	return exercise, nil
}