package courses

import "stringtheory/shared"


// loadAllCourses is responsible for asking
// the database for all courses.
// This function returns a slice of courses
// and an error
func loadAllCourses() ([]shared.Course, error) {
	c, err := sm.DataStore().getAllCourses()
	if err != nil {
		return c, err
	}
	return c, nil
}

// loadCourse is responsible for asking the
// database for a single course.
// This function accepts a course id and
// returns a course and an error
func loadCourse(cI string) (shared.Course, error) {
	c, err := sm.DataStore().getCourse(cI)
	return c, err
}