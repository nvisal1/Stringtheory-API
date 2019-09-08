package lessons

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"Stringtheory-API/shared"
)

type moduleMongoDataStore struct {
	db *mongo.Database
}

func (mmds moduleMongoDataStore) getCourseLessons(cI string) ([]shared.Lesson, error) {
	var result []shared.Lesson

	filter := bson.D{{
		"courseId",
		cI,
	}}

	cur, err := mmds.db.Collection("Lessons").Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		log.Fatal(err)
	}

	err = cur.All(context.TODO(), &result)
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

