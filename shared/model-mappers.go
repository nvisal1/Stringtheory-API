package shared

import (
	"errors"
	"os"
)

func NewCourse(ID string, name string, description string) (Course, error) {
	lessonsURI, err := generateLessonURI(ID)
	generatedCourse := Course{ ID, name, description, lessonsURI }
	if err != nil {
		return generatedCourse, err
	}
	return generatedCourse, nil
}

func generateLessonURI(courseID string) (string, error) {
	serviceDomain, exists := os.LookupEnv("SERVICE_DOMAIN")
	if exists {
		return serviceDomain + "/courses/" + courseID  + "/lessons", nil
	}
	return "",  errors.New("service domain not set")
}

func NewLesson(
	ID string,
	name string,
	order int32,
	description string,
	courseID string,
	hasNext bool,
) (Lesson, error) {
	exerciseURI, err := generateExerciseURI(courseID, ID)
	generatedLesson := Lesson{ID, name, order, description, courseID, hasNext, exerciseURI}
	if err != nil {
		return generatedLesson, err
	}
	return generatedLesson, nil
}

func generateExerciseURI(courseID string, lessonID string) (string, error) {
	serviceDomain, exists := os.LookupEnv("SERVICE_DOMAIN")
	if exists {
		return serviceDomain + "/courses/" + courseID  + "/lessons/" + lessonID + "/exercises", nil
	}
	return "",  errors.New("service domain not set")
}

func NewExercise(
	ID string,
	name string,
	description string,
	order int32,
	lessonId string,
	notes []string,
	hasNext bool,
) Exercise {
	generatedExercise := Exercise{
		ID:          ID,
		Name:        name,
		Order:       order,
		Notes:       notes,
		Description: description,
		LessonId:    lessonId,
		HasNext:     hasNext,
	}

	return generatedExercise
}