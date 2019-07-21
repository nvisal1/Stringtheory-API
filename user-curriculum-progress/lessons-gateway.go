package user_curriculum_progress

import (
	"stringtheory/lessons"
	"stringtheory/shared"
)

func getLesson(lI string) (shared.Lesson, error) {
	l, err := lessons.GetInstance().LoadLesson(lI)
	return l, err
}
