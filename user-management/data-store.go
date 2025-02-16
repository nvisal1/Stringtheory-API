package user_management

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

func (mmds moduleMongoDataStore) getUser(un string) (shared.User, error) {
	var result shared.User

	filter := bson.D{{"username", un}}

	err := mmds.db.Collection("Users").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		log.Fatal(err)
	}
	return result, nil
}

func (mmds moduleMongoDataStore) createUser(u shared.User) error {
	_, err := mmds.db.Collection("Users").InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	return nil
}
