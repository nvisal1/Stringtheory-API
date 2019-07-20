package exercises

import (
	"log"
	"sync"
)

type exercisesAdapter struct {}

var instance exercisesAdapter
var once sync.Once

func (ea exercisesAdapter) LoadLessonExercises(lI string) ([]exercise, error) {
	e, err := loadLessonExercises(lI)
	return e, err
}

func (ea exercisesAdapter) LoadExercise(eI string) (exercise, error) {
	e, err := loadExercise(eI)
	return e, err
}

func OpenInternalAdapter() {
	once.Do(func() {
		instance = exercisesAdapter{}
	})
}

func GetInstance() exercisesAdapter {
	if instance.LoadExercise == nil {
		log.Fatal("Lessons Adapter has not been created")
	}
	return instance
}
