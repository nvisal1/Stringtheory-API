package courses

type stubMongoDataStore struct {}

func (mmds stubMongoDataStore) getAllCourses() ([]course, error) {
	s := make([]string, 3)
	courses := make([]course, 1)
	sum := summary {
		"test",
		"test",
		"test",
	}
	u := course{
		sum,
		s,
	}
	res := append(courses, u)
	return res, nil
}