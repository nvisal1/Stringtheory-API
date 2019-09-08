package courses

import (
	"Stringtheory-API/shared"
	"sync"
)

type coursesAdapter struct {}

var instance coursesAdapter
var once sync.Once

// InitializeAdapter is required by the shared/Adapters
// interface. The function is responsible for
// creating a singleton instance of the
// coursesAdapter type
func (ca coursesAdapter) InitializeAdapter() {
	once.Do(func() {
		instance = coursesAdapter{}
	})
}

// LoadAllCourses is responsible for
// proxying incoming requests to the
// loadAllCourses interactor function
func (ca coursesAdapter) LoadAllCourses() ([]shared.Course, error) {
	c, err := loadAllCourses()
	return c, err
}

// LoadCourse is responsible for
// proxying incoming requests to
// the loadCourse interactor
// function
func (ca coursesAdapter) LoadCourse(cI string) (shared.Course, error) {
	c, err := loadCourse(cI)
	return c, err
}

// GetInstance returns the singleton instance
// that is created in the InitializeAdapter
// function. If an instance does not exist,
// the function throws an error and the service
// exits.
func GetInstance() coursesAdapter {
	return instance
}