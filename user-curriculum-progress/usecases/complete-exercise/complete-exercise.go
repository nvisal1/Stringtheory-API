package complete_exercise

import (
	"Stringtheory-API/shared"
	. "Stringtheory-API/user-curriculum-progress/service-module"
	. "Stringtheory-API/user-curriculum-progress/types"
	"errors"
	"fmt"
	"time"
)

func CompleteExercise(requester *shared.User, username string, exerciseID string, score float64) error {
	// check that requester is completed their own exercise
	if requester.Username != username {
		return errors.New("A user is not allowed to completed an exercise for another user")
	}

	// get user to verify existence
	_, err := UserCurriculumProgressModule.ServiceCommunicator.GetUser(username)
	if err != nil {
		return err
	}

	// get exercise to verify existence
	exercise, err := UserCurriculumProgressModule.ServiceCommunicator.GetExercise(exerciseID)
	if err != nil {
		return err
	}

	// save completed exercise
	completedExercise := NewCompletedExercise(username, exerciseID, time.Now(), score)
	err = UserCurriculumProgressModule.DataStore.CreateCompletedExercise(completedExercise)
	if err != nil {
		return err
	}

	// send completed exercise message
	message := createCompletedExerciseMessage(requester, completedExercise, exercise.LessonId)
	UserCurriculumProgressModule.MessageStore.SendMessage(message)

	return nil
}

func createCompletedExerciseMessage(requester *shared.User, completedExercise *CompletedExercise, lessonID string) *Message {
	messageBody := fmt.Sprintf("User %s completed exercise %s", requester.Username, completedExercise.ExerciseId)
	messageAttributes := make(map[string]interface{})
	messageAttributes["title"] = "exercise completed"
	messageAttributes["requester"] = requester
	messageAttributes["exercise"] = completedExercise
	messageAttributes["lessonID"] = lessonID
	message := NewMessage(messageBody, messageAttributes)
	return message
}