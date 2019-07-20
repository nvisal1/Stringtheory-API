package user_curriculum_progress

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type moduleMongoDataStore struct {
	db *mongo.Database
}

func (mmds moduleMongoDataStore) createCompletedCourse(cC completedCourse) error {
	_, err := mmds.db.Collection("Completed-Courses").InsertOne(context.TODO(), cC)
	return err
}

func (mmds moduleMongoDataStore) createCompletedLesson(cL completedLesson) error {
	_, err := mmds.db.Collection("Completed-Lessons").InsertOne(context.TODO(), cL)
	return err
}

func (mmds moduleMongoDataStore) createCompletedExercise(cE completedExercise) error {
	_, err := mmds.db.Collection("Completed-Exercises").InsertOne(context.TODO(), cE)
	return err
}