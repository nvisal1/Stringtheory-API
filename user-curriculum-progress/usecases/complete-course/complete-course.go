package complete_course

import (
	. "Stringtheory-API/user-curriculum-progress/service-module"
	. "Stringtheory-API/user-curriculum-progress/types"
	"time"
)

func CompleteCourse(username string, courseID string) error {
	// get user to verify existence
	_, err := UserCurriculumProgressModule.ServiceCommunicator.GetUser(username)
	if err != nil {
		return err
	}

	// get course to verify existence
	_, err = UserCurriculumProgressModule.ServiceCommunicator.GetCourse(courseID)
	if err != nil {
		return err
	}

	// make sure that the user hasn't completed the course already
    _, err = UserCurriculumProgressModule.Datastore.GetCompletedCourse(username)
    if err != nil {
    	return err
	}

	// make sure that the user is eligible to complete the specified course
	err = checkCourseProgress(username, courseID)
	if err != nil {
		return err
	}

	completedCourse := NewCompletedCourse(username, courseID, time.Now())

	err = UserCurriculumProgressModule.Datastore.CreateCompletedCourse(completedCourse)
	if err != nil {
		return err
	}

	return err
}


// checkCourseProgress checks if the user completed all of the lessons
// in a given course. The function returns an error if the user
// is not eligible to complete the specified course
func checkCourseProgress(username string, courseID string) error {

}