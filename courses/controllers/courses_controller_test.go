package controllers

import (
	"Stringtheory-API/courses"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	courses.InitializeModule()
	os.Exit(m.Run())
}


func TestAddWithGoConvey(t *testing.T) {
	Convey("Adding two numbers", t, func() {
		x := 1
		y := 2

		Convey("should produce the expected result", func() {
			So(Add(x, y), ShouldEqual, 3)
		})
	})
}

