package load_course_lessons

import (
	. "Stringtheory-API/lessons/drivers/data-store"
	. "Stringtheory-API/lessons/drivers/service-communication"
	. "Stringtheory-API/lessons/service-module"
	. "Stringtheory-API/shared"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("SERVICE_ENVIRONMENT", "test")
	NewLessonsModule(nil, StubMongoDataStore{}, StubServiceCommunicator{} )
}

func TestLoadCourseLessons(t *testing.T) {
	Convey("When LoadCourseLessons is called", t, func() {
		courseLessons, err := LoadCourseLessons("test_courseID")

		Convey("and no driver error occurs", func() {
			expectedResponse := []Lesson{{
				"test",
				"test",
				1,
				"test",
				"test",
				true,
				"test",
			}}

			Convey("a slice of Lesson should be returned without error", func() {
				So(courseLessons, ShouldResemble, expectedResponse)
				So(err, ShouldBeNil)
			})
		})
	})
}