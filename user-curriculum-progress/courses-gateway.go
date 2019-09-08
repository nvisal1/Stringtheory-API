package user_curriculum_progress

import (
	"Stringtheory-API/courses"
	"Stringtheory-API/shared"
)

func getCourse(cI string) (shared.Course, error) {
	c, err := courses.GetInstance().LoadCourse(cI)
	return c, err
}