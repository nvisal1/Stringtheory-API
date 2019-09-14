package lessons

import "Stringtheory-API/shared"

func loadCourseLessons(cI string) ([]shared.Lesson, error) {
	l, err := sm.DataStore().getCourseLessons(cI)
	if err != nil {
		return l, err
	}
	return l, nil
}

func loadLesson(lI string) (shared.Lesson, error) {
	l, err := sm.DataStore().getLesson(lI)
	return l, err
}
