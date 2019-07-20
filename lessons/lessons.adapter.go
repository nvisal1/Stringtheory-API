package lessons

import (
	"log"
	"sync"
)

type lessonsAdapter struct {}

var instance lessonsAdapter
var once sync.Once

func (la lessonsAdapter) LoadCourseLessons(cI string) ([]lesson, error) {
	l, err := loadCourseLessons(cI)
	return l, err
}

func (la lessonsAdapter) LoadLesson(lI string) (lesson, error) {
	l, err := loadLesson(lI)
	return l, err
}

func OpenInternalAdapter() {
	once.Do(func() {
		instance = lessonsAdapter{}
	})
}

func GetInstance() lessonsAdapter {
	if instance.LoadCourseLessons == nil {
		log.Fatal("Lessons Adapter has not been created")
	}
	return instance
}