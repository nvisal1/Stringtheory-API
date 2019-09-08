package user_curriculum_progress

import (
	"Stringtheory-API/lessons"
	"Stringtheory-API/shared"
)

func getLesson(lI string) (shared.Lesson, error) {
	l, err := lessons.GetInstance().LoadLesson(lI)
	return l, err
}
