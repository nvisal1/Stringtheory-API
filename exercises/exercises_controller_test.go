package exercises

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitializeModule()
	os.Exit(m.Run())
}

func TestLoadLessonExercises_With_Result(t *testing.T) {
	_, err := loadLessonExercises("test")

	if err != nil {
		t.Error("loadLessonExercises returned an unexpected error")
	}
}

func TestLoadLessonExercises_No_Result(t *testing.T) {
	_, err := loadLessonExercises("")

	if err == nil {
		t.Error("loadLessonExercises did not return the expected error")
	}
}

func TestLoadExercise_With_Result(t *testing.T) {
	_, err := loadExercise("test")

	if err != nil {
		t.Error("loadExercise returned an unexpected error")
	}
}

func TestLoadExerecise_No_Result(t *testing.T) {
	_, err := loadExercise("")

	if err == nil {
		t.Error("loadExercise did not return the expected error")
	}
}