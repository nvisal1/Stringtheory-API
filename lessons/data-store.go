package lessons

import (
	"Stringtheory-API/shared"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type moduleMongoDataStore struct {
	db *mongo.Database
}

func (mmds moduleMongoDataStore) getCourseLessons(cI string) ([]shared.Lesson, error) {
	var result []shared.Lesson

	filter := bson.D{{
		"CourseId",
		cI,
	}}

	cur, err := mmds.db.Collection("Lessons").Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		elem := &bson.D{}
		if err = cur.Decode(elem); err != nil {
			return result, err
		}
		m := elem.Map()
		l := shared.Lesson{
			ID: m["_id"].(primitive.ObjectID).Hex(),
			Name:  m["Name"].(string),
			Order: m["Order"].(int32),
			Description: m["Description"].(string),
			CourseId: m["CourseId"].(string),
			HasNext: m["HasNext"].(bool),
		}
		result = append(result, l)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func (mmds moduleMongoDataStore) getLesson(lI string) (shared.Lesson, error) {
	var result shared.Lesson

	filter := bson.D{{"_id", lI}}

	err := mmds.db.Collection("Lessons").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		log.Fatal(err)
	}
	return result, nil
}

