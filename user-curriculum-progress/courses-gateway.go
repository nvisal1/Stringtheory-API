package user_curriculum_progress

import (
	. "Stringtheory-API/courses/usecases/load-course"
	"Stringtheory-API/shared"
)

func getCourse(courseID string) (shared.Course, error) {
	c, err := LoadCourse(courseID)
	return c, err
}