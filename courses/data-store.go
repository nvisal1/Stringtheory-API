package courses

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"stringtheory/shared"
)

type moduleMonogDataStore struct {
	db *mongo.Database
}

func (mmds moduleMonogDataStore) getAllCourses() ([]shared.Course, error) {
	var result []shared.Course

	filter := bson.D{{}}

	cur, err := mmds.db.Collection("Courses").Find(context.TODO(), filter)
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

func (mmds moduleMonogDataStore) getCourse(cI string) (shared.Course, error) {
	var result shared.Course

	filter := bson.D{{"_id", cI}}

	err := mmds.db.Collection("Courses").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		log.Fatal(err)
	}
	return result, nil
}