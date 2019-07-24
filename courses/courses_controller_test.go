package courses

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitializeModule()
	os.Exit(m.Run())
}

func TestLoadAllCourses(t *testing.T) {
	_, err := loadAllCourses()

	if err != nil {
		t.Error("loadAllCourses returned an unexpected error")
	}
}

func TestLoadCourse_With_Result(t *testing.T) {
	_, err := loadCourse("test")

	if err != nil {
		t.Error("loadAllCourses returned an unexpected error")
	}
}

func TestLoadCourse_No_Result(t *testing.T) {
	_, err := loadCourse("")

	if err == nil {
		t.Error("loadAllCourses did not return the expected error")
	}
}

