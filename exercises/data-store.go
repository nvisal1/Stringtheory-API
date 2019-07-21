package exercises

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"stringtheory/shared"
)

type moduleMongoDataStore struct {
	db *mongo.Database
}

func (mmds moduleMongoDataStore) getLessonExercises(lI string) ([]shared.Exercise, error) {
	var result []shared.Exercise

	filter := bson.D{{
		"_id",
		lI,
	}}

	cur, err := mmds.db.Collection("Exercises").Find(context.TODO(), filter)
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

func (mmds moduleMongoDataStore) getExercise(eI string) (shared.Exercise, error) {
	var result shared.Exercise

	filter := bson.D{{"_id", eI}}

	err := mmds.db.Collection("Exercises").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		log.Fatal(err)
	}
	return result, nil
}