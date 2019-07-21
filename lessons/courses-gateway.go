package lessons

import (
	"stringtheory/courses"
	"stringtheory/shared"
)

func getCourse(cI string) (shared.Course, error) {
	c, err := courses.GetInstance().LoadCourse(cI)
	return c, err
}