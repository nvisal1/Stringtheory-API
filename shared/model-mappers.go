package shared

import (
	"errors"
	"os"
)

func GenerateCourseFromStoredData(ID string, name string, description string) (Course, error) {
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