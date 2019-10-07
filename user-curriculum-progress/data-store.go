package user_curriculum_progress

import (
	"Stringtheory-API/shared"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
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

func (mmds moduleMongoDataStore) getCompletedLesson(uN string) (completedLesson, error) {
	var result []completedLesson

	filter := bson.D{{"username", uN}}

	cur, err := mmds.db.Collection("Completed-Lessons").Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result[0], err
		}
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		elem := &bson.D{}
		if err = cur.Decode(elem); err != nil {
			return result[0], err
		}
		m := elem.Map()
		c := completedLesson{
			ID:         m["_id"].(primitive.ObjectID).Hex(),
			Username: m["Username"].(string),
			LessonId: m["LessonId"].(string),
			Date: m["Date"].(time.Time),
		}
		result = append(result, c)
	}
	return result[0], nil
}

func (mmds moduleMongoDataStore) getCompletedCourse(uN string) (completedCourse, error) {
	var result []completedCourse

	filter := bson.D{{"username", uN}}

	cur, err := mmds.db.Collection("Completed-Lessons").Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result[0], err
		}
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		elem := &bson.D{}
		if err = cur.Decode(elem); err != nil {
			return result[0], err
		}
		m := elem.Map()
		c := completedCourse{
			ID:         m["_id"].(primitive.ObjectID).Hex(),
			Username: m["Username"].(string),
			CourseId: m["CourseId"].(string),
			Date: m["Date"].(time.Time),
		}
		result = append(result, c)
	}
	return result[0], nil
}