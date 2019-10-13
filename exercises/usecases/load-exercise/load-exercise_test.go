package load_exercise

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

func TestLoadExercise(t *testing.T) {
	Convey("When LoadExercise is called", t, func() {
		exercise, err := LoadExercise("test_exerciseID")

		Convey("and no driver error occurs", func() {
			expectedResponse := shared.Exercise{
				"test",
				"test",
				0,
				make([]string, 3),
				"test",
				"test",
				false,
			}

			Convey("a single Exercise should be returned without error", func() {
				So(exercise, ShouldResemble, expectedResponse)
				So(err, ShouldBeNil)
			})
		})
	})
}

