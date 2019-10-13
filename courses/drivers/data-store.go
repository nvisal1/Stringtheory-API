package drivers

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

// getAllCourses is responsible for finding and returning all
// documents in the Courses collection.
func (datastore ModuleMongoDataStore) GetAllCourses() ([]shared.Course, error) {
	var result []shared.Course

	filter := bson.D{{}}

	cur, err := datastore.DB.Collection("Courses").Find(context.TODO(), filter)
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
		ID := m["_id"].(primitive.ObjectID).Hex()
		name := m["Name"].(string)
		description :=  m["Description"].(string)
		c, err := shared.NewCourse(ID, name, description)
		if err != nil {
			return result, err
		}
		result = append(result, c)
	}
	return result, nil
}

// getCourse is responsible for finding and returning
// a single document in the Courses collection
// that matches the given course id. If no result
// is found, the function returns an error.
func (datastore ModuleMongoDataStore) GetCourse(cI string) (shared.Course, error) {
	var result shared.Course

	filter := bson.D{{"_id", cI}}

	err := datastore.DB.Collection("Courses").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		log.Fatal(err)
	}
	return result, nil
}