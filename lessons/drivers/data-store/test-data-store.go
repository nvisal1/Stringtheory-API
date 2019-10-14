package data_store

import "Stringtheory-API/shared"

type StubMongoDataStore struct {}

func (datastore StubMongoDataStore) QueryCourseLessons(courseID string) ([]shared.Lesson, error) {
	lessons := make([]shared.Lesson, 3)
	lesson := shared.Lesson {
		"test",
		"test",
		1,
		"test",
		"test",
		true,
		"test",
	}
	return append(lessons, lesson), nil
}

func (datastore StubMongoDataStore) QueryLesson(lessonID string) (shared.Lesson, error) {
	lesson := shared.Lesson {
		"test",
		"test",
		1,
		"test",
		"test",
		true,
		"test",
	}

	return lesson, nil
}