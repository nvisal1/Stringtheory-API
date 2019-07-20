package exercises

type stubMongoDataStore struct {}

func (mmds stubMongoDataStore) getLessonExercises(lE []string) ([]exercise, error) {
	es := make([]exercise, 3)
	e := exercise {
		"test",
		"test",
		make([]string, 3),
		"test",
	}
	return append(es, e), nil
}
