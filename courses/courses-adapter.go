package courses

import (
	"log"
	"sync"
)

type coursesAdapter struct {}

var instance coursesAdapter
var once sync.Once

func (ca coursesAdapter) LoadAllCourses() ([]course, error) {
	c, err := loadAllCourses()
	return c, err
}

func (ca coursesAdapter) LoadCourse(cI string) (course, error) {
	c, err := loadCourse(cI)
	return c, err
}

func OpenInternalAdapter() {
	once.Do(func() {
		instance = coursesAdapter{}
	})
}

func GetInstance() coursesAdapter {
	if instance.LoadAllCourses == nil {
		log.Fatal("Lessons Adapter has not been created")
	}
	return instance
}