package complete_lesson

import (
	"Stringtheory-API/shared"
	. "Stringtheory-API/user-curriculum-progress/service-module"
	. "Stringtheory-API/user-curriculum-progress/types"
	"errors"
	"fmt"
	"time"
)

func CompleteLesson(requester *shared.User, lessonID string) error {
	// get requester to verify existence
	_, err := UserCurriculumProgressModule.ServiceCommunicator.GetUser(requester.Username)
	if err != nil {
		return err
	}

	// get lesson to verify existence
	lesson, err := UserCurriculumProgressModule.ServiceCommunicator.GetLesson(lessonID)
	if err != nil {
		return err
	}

	// make sure that the user hasn't completed the lesson already
	_, err = UserCurriculumProgressModule.DataStore.GetCompletedLesson(requester.Username)
	if err == nil {
		return errors.New("user already completed the specified Lesson")
	}

	completedLesson := NewCompletedLesson(requester.Username, lessonID, time.Now())

	err = UserCurriculumProgressModule.DataStore.CreateCompletedLesson(completedLesson)
	if err != nil {
		return err
	}

	// send completed lesson message
	message := createCompletedLessonMessage(requester, completedLesson, lesson.CourseId)
	UserCurriculumProgressModule.MessageStore.SendMessage(message)

	return nil
}

func createCompletedLessonMessage(requester *shared.User, completedLesson *CompletedLesson, courseID string) *Message {
	messageBody := fmt.Sprintf("User %s completed lesson %s", requester.Username, completedLesson.LessonId)
	messageAttributes := make(map[string]interface{})
	messageAttributes["title"] = "lesson completed"
	messageAttributes["requester"] = requester
	messageAttributes["lesson"] = completedLesson
	messageAttributes["courseID"] = courseID
	message := NewMessage(messageBody, messageAttributes)
	return message
}