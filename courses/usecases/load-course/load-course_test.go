package load_course

import (
	"Stringtheory-API/courses/drivers"
	. "Stringtheory-API/courses/service-module"
	"Stringtheory-API/shared"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("SERVICE_ENVIRONMENT", "test")
	NewCoursesModule(nil, drivers.StubMongoDataStore{})
}

func TestLoadCourse(t *testing.T) {
	Convey("When LoadCourse is called", t, func() {
		courses, err := LoadCourse("test_courseID")

		Convey("And no driver error occur", func() {
			expectedResponse := shared.Course{"test", "test", "test", "test"}

			Convey("a Course should be returned without error", func() {
				So(courses, ShouldResemble, expectedResponse)
				So(err, ShouldBeNil)
			})
		})
	})
}
