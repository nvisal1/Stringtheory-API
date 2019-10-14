package load_course_lessons

import (
	. "Stringtheory-API/lessons/service-module"
	"Stringtheory-API/shared"
)

func LoadCourseLessons(courseID string) ([]shared.Lesson, error) {
	courseLessons, err := LessonsModule.Datastore.QueryCourseLessons(courseID)
	return courseLessons, err
}