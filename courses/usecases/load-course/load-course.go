package load_course

import (
	service_module "Stringtheory-API/courses/service-module"
	"Stringtheory-API/shared"
)

// LoadCourse is responsible for asking
// the database for single specified course.
// This function returns a course
// and an error
func LoadCourse(courseID string) (shared.Course, error) {
	course, err := service_module.CoursesModule.Datastore.GetCourse(courseID)
	if err != nil {
		return course, err
	}
	return course, nil
}