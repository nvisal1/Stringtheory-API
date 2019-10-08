package courses

import (
	"errors"
	"os"
)


// loadAllCourses is responsible for asking
// the database for all courses.
// This function returns a slice of courses
// and an error
func loadAllCourses() ([]shared.Course, error) {
	c, err := sm.ds.getAllCourses()
	if err != nil {
		return c, err
	}

	withLessonURIs := c[:0]
	for _, course := range c {
		course.LessonsURI, err = generateLessonURI(course.ID)
		if err != nil {
			return c, err
		}
		withLessonURIs = append(withLessonURIs, course)
	}
	return withLessonURIs, nil
}

// loadCourse is responsible for asking the
// database for a single course.
// This function accepts a course id and
// returns a course and an error
func loadCourse(cI string) (shared.Course, error) {
	c, err := sm.ds.getCourse(cI)
	return c, err
}

func generateLessonURI(courseID string) (string, error) {
	serviceDomain, exists := os.LookupEnv("SERVICE_DOMAIN")
	if exists {
		return serviceDomain + "/courses/" + courseID  + "/lessons", nil
	}
	return "",  errors.New("service domain not set")
}