package load_all_courses

import (
	service_module "Stringtheory-API/courses/service-module"
	"Stringtheory-API/shared"
)

// LoadAllCourses is responsible for asking
// the database for all courses.
// This function returns a slice of courses
// and an error
func LoadAllCourses() ([]shared.Course, error) {
	courses, err := service_module.CoursesModule.Datastore.GetAllCourses()
	if err != nil {
		return courses, err
	}
	return courses, nil
}