package lessons

import "stringtheory/shared"

type dataStore interface {
	getCourseLessons(cI string) ([]shared.Lesson, error)
	getLesson(lI string) (shared.Lesson, error)
}
