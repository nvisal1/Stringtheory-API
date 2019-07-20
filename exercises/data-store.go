package exercises

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type moduleMongoDataStore struct {
	db *mongo.Database
}

func (mmds moduleMongoDataStore) getLessonExercises(lE []string) ([]exercise, error) {
	var result []exercise

	filter := bson.D{{
		"_id",
		bson.D{{
			"$in",
			bson.A{lE},
		}},
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
