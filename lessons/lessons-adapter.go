package lessons

import (
	"Stringtheory-API/shared"
	"sync"
)

type lessonsAdapter struct {}

var instance lessonsAdapter
var once sync.Once

func (la lessonsAdapter) InitializeAdapter() {
	once.Do(func() {
		instance = lessonsAdapter{}
	})
}

func (la lessonsAdapter) LoadCourseLessons(cI string) ([]shared.Lesson, error) {
	l, err := loadCourseLessons(cI)
	return l, err
}

func (la lessonsAdapter) LoadLesson(lI string) (shared.Lesson, error) {
	l, err := loadLesson(lI)
	return l, err
}

func GetInstance() lessonsAdapter {
	return instance
}