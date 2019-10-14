package load_lesson

import (
	. "Stringtheory-API/lessons/service-module"
	"Stringtheory-API/shared"
)

func LoadLesson(lessonID string) (shared.Lesson, error) {
	lesson, err := LessonsModule.Datastore.QueryLesson(lessonID)
	return lesson, err
}