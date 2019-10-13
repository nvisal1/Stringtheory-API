package load_lesson_exercises

import (
	"Stringtheory-API/exercises/drivers"
	. "Stringtheory-API/exercises/service-module"
	"Stringtheory-API/shared"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("SERVICE_ENVIRONMENT", "test")
	NewExercisesModule(nil, drivers.StubMongoDataStore{})
}

func TestLoadLessonExercises(t *testing.T) {
	Convey("When LoadLessonExercises is called", t, func() {
		exercises, err := LoadLessonExercises("test_lessonID")

		Convey("and no driver error occurs", func() {
			expectedResponse := []shared.Exercise{{
				"test",
				"test",
				0,
				make([]string, 3),
				"test",
				"test",
				false,
			}}

			Convey("a slice of Exercises should be returned without error", func() {
				So(exercises, ShouldResemble, expectedResponse)
				So(err, ShouldBeNil)
			})
		})
	})
}