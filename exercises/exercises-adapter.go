package exercises

import (
	"log"
	"stringtheory/shared"
	"sync"
)

type exercisesAdapter struct {}

var instance exercisesAdapter
var once sync.Once

func (ea exercisesAdapter) InitializeAdapter() {
	once.Do(func() {
		instance = exercisesAdapter{}
	})
}

func (ea exercisesAdapter) LoadLessonExercises(lI string) ([]shared.Exercise, error) {
	e, err := loadLessonExercises(lI)
	return e, err
}

func (ea exercisesAdapter) LoadExercise(eI string) (shared.Exercise, error) {
	e, err := loadExercise(eI)
	return e, err
}

func GetInstance() exercisesAdapter {
	if instance.LoadExercise == nil {
		log.Fatal("Lessons Adapter has not been created")
	}
	return instance
}
