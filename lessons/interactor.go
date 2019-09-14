package lessons

import (
	"Stringtheory-API/shared"
	"errors"
	"os"
)

func loadCourseLessons(cI string) ([]shared.Lesson, error) {
	l, err := sm.DataStore().getCourseLessons(cI)
	if err != nil {
		return l, err
	}

	withExerciseURIs := l[:0]
	for _, lesson := range l {
		lesson.ExercisesURI, err = generateExerciseURI(lesson.CourseId, lesson.ID)
		if err != nil {
			return l, err
		}
		withExerciseURIs = append(withExerciseURIs, lesson)
	}

	return withExerciseURIs, nil
}

func loadLesson(lI string) (shared.Lesson, error) {
	l, err := sm.DataStore().getLesson(lI)
	return l, err
}

func generateExerciseURI(courseID string, lessonID string) (string, error) {
	serviceDomain, exists := os.LookupEnv("SERVICE_DOMAIN")
	if exists {
		return serviceDomain + "/courses/" + courseID  + "/lessons/" + lessonID + "/exercises", nil
	}
	return "",  errors.New("service domain not set")
}