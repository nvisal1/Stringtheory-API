package load_all_courses

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

func TestLoadAllCourses(t *testing.T) {
	Convey("When LoadAllCourses is called", t, func() {
		courses, err := LoadAllCourses()

		Convey("And no driver error occur", func() {
			expectedResponse := []shared.Course{{"test", "test", "test", "test"}}

			Convey("a slice of Course should be returned without error", func() {
				So(courses, ShouldResemble, expectedResponse)
				So(err, ShouldBeNil)
			})
		})
	})
}

