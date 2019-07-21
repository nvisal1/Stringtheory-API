package courses

import (
	"log"
	"stringtheory/shared"
	"sync"
)

type coursesAdapter struct {}

var instance coursesAdapter
var once sync.Once

func (ca coursesAdapter) InitializeAdapter() {
	once.Do(func() {
		instance = coursesAdapter{}
	})
}

func (ca coursesAdapter) LoadAllCourses() ([]shared.Course, error) {
	c, err := loadAllCourses()
	return c, err
}

func (ca coursesAdapter) LoadCourse(cI string) (shared.Course, error) {
	c, err := loadCourse(cI)
	return c, err
}

func GetInstance() coursesAdapter {
	if instance.LoadAllCourses == nil {
		log.Fatal("Lessons Adapter has not been created")
	}
	return instance
}