package data_store

import (
	"Stringtheory-API/shared"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type ModuleMongoDataStore struct {
	DB *mongo.Database
}

func (datastore ModuleMongoDataStore) QueryCourseLessons(courseID string) ([]shared.Lesson, error) {
	var result []shared.Lesson

	filter := bson.D{{
		"CourseId",
		courseID,
	}}

	cursor, err := datastore.DB.Collection("Lessons").Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		lesson, err := datastore.deserializeDocumentFromCursor(cursor)
		if err != nil {
			return result, err
		}
		result = append(result, lesson)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func (datastore ModuleMongoDataStore) QueryLesson(lessonID string) (shared.Lesson, error) {
	var result shared.Lesson

	filter := bson.D{{"_id", lessonID}}

	err := datastore.DB.Collection("Lessons").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		log.Fatal(err)
	}
	return result, nil
}

func (datastore ModuleMongoDataStore) deserializeDocumentFromCursor(cursor *mongo.Cursor) (shared.Lesson, error) {
	element := &bson.D{}
	if err := cursor.Decode(element); err != nil {
		return shared.Lesson{}, err
	}
	elementMap := element.Map()
	ID := elementMap["_id"].(primitive.ObjectID).Hex()
	name := elementMap["Name"].(string)
	order := elementMap["Order"].(int32)
	description := elementMap["Description"].(string)
	courseID := elementMap["CourseId"].(string)
	hasNext := elementMap["HasNext"].(bool)

	lesson, err := shared.NewLesson(ID, name, order, description, courseID, hasNext)
	if err != nil {
		return shared.Lesson{}, err
	}
	return lesson, nil
}

