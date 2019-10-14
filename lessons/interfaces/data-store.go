package interfaces

import "Stringtheory-API/shared"

type Datastore interface {
	QueryCourseLessons(cI string) ([]shared.Lesson, error)
	QueryLesson(lI string) (shared.Lesson, error)
}
