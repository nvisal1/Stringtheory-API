package courses

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

// getAllCourses is responsible for finding and returning all
// documents in the Courses collection.
func (mmds moduleMongoDataStore) getAllCourses() ([]shared.Course, error) {
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

// getCourse is responsible for finding and returning
// a single document in the Courses collection
// that matches the given course id. If no result
// is found, the function returns an error.
func (mmds moduleMongoDataStore) getCourse(cI string) (shared.Course, error) {
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