package user_curriculum_progress

import (
	"Stringtheory-API/courses/adapters/internal"
	"Stringtheory-API/shared"
)

func getCourse(cI string) (shared.Course, error) {
	c, err := internal.GetInstance().LoadCourse(cI)
	return c, err
}