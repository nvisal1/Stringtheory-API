package user_curriculum_progress

import (
	"stringtheory/courses/adapters"
	"stringtheory/shared"
)

func getCourse(cI string) (shared.Course, error) {
	c, err := adapters.GetInstance().LoadCourse(cI)
	return c, err
}