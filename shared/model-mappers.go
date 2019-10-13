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