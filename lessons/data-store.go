package lessons

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type moduleMongoDataStore struct {
	db *mongo.Database
}

func (mmds moduleMongoDataStore) getCourseLessons(cL []string) ([]lesson, error) {
	var result []lesson

	filter := bson.D{{
		"_id",
		bson.D{{
			"$in",
			bson.A{cL},
		}},
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