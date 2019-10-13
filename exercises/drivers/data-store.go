package drivers

import (
	. "Stringtheory-API/shared"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type ModuleMongoDataStore struct {
	DB *mongo.Database
}

func (datastore ModuleMongoDataStore) QueryLessonExercises(lessonID string) ([]Exercise, error) {
	var result []Exercise

	filter := bson.D{{
		"LessonId",
		lessonID,
	}}

	cursor, err := datastore.DB.Collection("Exercises").Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, NewServiceError("The specified exercise was not found.", ErrorTypes.NotFound)
		}
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		exercise, err := datastore.deserializeDocumentFromCursor(cursor)
		if err != nil {
			return result, err
		}
		result = append(result, exercise)
	}
	return result, nil
}

func (datastore ModuleMongoDataStore) QueryExercise(eI string) (Exercise, error) {
	var result Exercise

	filter := bson.D{{"_id", eI}}

	err := datastore.DB.Collection("Exercises").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, NewServiceError("The specified exercise was not found.", ErrorTypes.NotFound)
		}
		log.Fatal(err)
	}
	return result, nil
}

func (datastore ModuleMongoDataStore) deserializeDocumentFromCursor(cur *mongo.Cursor) (Exercise, error) {
	element := &bson.D{}
	if err := cur.Decode(element); err != nil {
		return Exercise{}, err
	}
	elementMap := element.Map()
	notes := make([]string, len(elementMap["Notes"].(primitive.A)))
	for i, v := range elementMap["Notes"].(primitive.A) {
		notes[i] = fmt.Sprint(v)
	}
	ID := elementMap["_id"].(primitive.ObjectID).Hex()
	name := elementMap["Name"].(string)
	description := elementMap["Description"].(string)
	order := elementMap["Order"].(int32)
	lessonID := elementMap["LessonId"].(string)
	hasNext := elementMap["HasNext"].(bool)

	exercise := NewExercise(ID, name, description, order, lessonID, notes, hasNext)
	return exercise, nil
}