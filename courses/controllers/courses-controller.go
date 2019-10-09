package controllers

import (
	courses2 "Stringtheory-API/courses"
	"Stringtheory-API/shared"
	"errors"
	"os"
)


// loadAllCourses is responsible for asking
// the database for all courses.
// This function returns a slice of courses
// and an error
func loadAllCourses() ([]shared.Course, error) {
	courses, err := courses2.getAllCourses()
	if err != nil {
		return courses, err
	}

	coursesWithLessonURIs, err := handleAppendingLessonURIs(c, err)
	if err != nil {
		return courses, err
	}
	return coursesWithLessonURIs, nil
}

func handleAppendingLessonURIs(originalCourses []shared.Course, err error) ([]shared.Course, error) {
	coursesWithLessonURIs := originalCourses[:0]
	for _, course := range originalCourses {
		course.LessonsURI, err = generateLessonURI(course.ID)
		if err != nil {
			return nil, err
		}
		coursesWithLessonURIs = append(coursesWithLessonURIs, course)
	}
	return coursesWithLessonURIs, nil
}

func generateLessonURI(courseID string) (string, error) {
	serviceDomain, exists := os.LookupEnv("SERVICE_DOMAIN")
	if exists {
		return serviceDomain + "/courses/" + courseID  + "/lessons", nil
	}
	return "",  errors.New("service domain not set")
}