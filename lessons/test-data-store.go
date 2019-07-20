package lessons

type stubMongoDataStore struct {}

func (mmds stubMongoDataStore) getCourseLessons(cL []string) ([]lesson, error) {
	es := make([]lesson, 3)
	e := lesson {
		"test",
		"test",
		make([]string, 3),
		"test",
	}
	return append(es, e), nil
}
