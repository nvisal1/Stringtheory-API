package data_store

import (
	. "Stringtheory-API/user-curriculum-progress/types"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type ModuleMongoDataStore struct {
	db *mongo.Database
}

func NewDatastore(databaseConnection *mongo.Database) *ModuleMongoDataStore {
	return &ModuleMongoDataStore{ db: databaseConnection}
}

func (datastore ModuleMongoDataStore) CreateCompletedCourse(completedCourse *CompletedCourse) error {
	_, err := datastore.db.Collection("Completed-Courses").InsertOne(context.TODO(), completedCourse)
	return err
}

func (datastore ModuleMongoDataStore) CreateCompletedLesson(completedLesson CompletedLesson) error {
	_, err := datastore.db.Collection("Completed-Lessons").InsertOne(context.TODO(), completedLesson)
	return err
}

func (datastore ModuleMongoDataStore) CreateCompletedExercise(completedExercise CompletedExercise) error {
	_, err := datastore.db.Collection("Completed-Exercises").InsertOne(context.TODO(), completedExercise)
	return err
}

func (datastore ModuleMongoDataStore) GetCompletedLesson(username string) (CompletedLesson, error) {
	var result []CompletedLesson

	filter := bson.D{{"username", username}}

	cur, err := datastore.db.Collection("Completed-Lessons").Find(context.TODO(), filter)
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
		c := CompletedLesson{
			Username: m["Username"].(string),
			LessonId: m["LessonId"].(string),
			Date: m["Date"].(time.Time),
		}
		result = append(result, c)
	}
	return result[0], nil
}

func (datastore ModuleMongoDataStore) GetCompletedCourse(username string) (CompletedCourse, error) {
	var result []CompletedCourse

	filter := bson.D{{"username", username}}

	cur, err := datastore.db.Collection("Completed-Courses").Find(context.TODO(), filter)
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
		c := CompletedCourse{
			Username: m["Username"].(string),
			CourseId: m["CourseId"].(string),
			Date: m["Date"].(time.Time),
		}
		result = append(result, c)
	}
	return result[0], nil
}