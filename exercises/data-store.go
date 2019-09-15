package exercises

import (
	"Stringtheory-API/shared"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type moduleMongoDataStore struct {
	db *mongo.Database
}

func (mmds moduleMongoDataStore) getLessonExercises(lI string) ([]shared.Exercise, error) {
	var result []shared.Exercise

	filter := bson.D{{
		"LessonId",
		lI,
	}}

	cur, err := mmds.db.Collection("Exercises").Find(context.TODO(), filter)
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
		ns := make([]string, len(m["Notes"].(primitive.A)))
		for i, v := range m["Notes"].(primitive.A) {
			ns[i] = fmt.Sprint(v)
		}
		e := shared.Exercise{
			ID: m["_id"].(primitive.ObjectID).Hex(),
			Name: m["Name"].(string),
			Description: m["Description"].(string),
			Order: m["Order"].(int32),
			LessonId: m["LessonId"].(string),
			Notes: ns,
			HasNext: m["HasNext"].(bool),
		}
		result = append(result, e)
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