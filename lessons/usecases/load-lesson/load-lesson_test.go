package load_lesson

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

func TestLoadLesson(t *testing.T) {
	Convey("When LoadLesson is called", t, func() {
		lesson, err := LoadLesson("test_lessonID")

		Convey("and no driver error occurs", func() {
			expectedResponse := Lesson{
				ID: "test",
				Name: "test",
				Order: 1,
				Description: "test",
				CourseId: "test",
				HasNext: true,
				ExercisesURI: "test",
			}

			Convey("a single Lesosn should be returned without error", func() {
				So(lesson, ShouldResemble, expectedResponse)
				So(err, ShouldBeNil)
			})
		})
	})
}