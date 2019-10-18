package complete_course

import (
	. "Stringtheory-API/user-curriculum-progress/service-module"
	. "Stringtheory-API/user-curriculum-progress/types"
	"time"
)

func CompleteCourse(username string, courseID string) error {
	// get course to verify existence
	_, err := UserCurriculumProgressModule.ServiceCommunicator.GetCourse(courseID)
	if err != nil {
		return err
	}
	// make sure that the user hasn't completed the course already
    _, err = UserCurriculumProgressModule.Datastore.GetCompletedCourse(username)
    if err != nil {
    	return err
	}

	completedCourse := NewCompletedCourse(username, courseID, time.Now())

	err = UserCurriculumProgressModule.Datastore.CreateCompletedCourse(completedCourse)
	return err
}

func checkLessonProgress() {

}