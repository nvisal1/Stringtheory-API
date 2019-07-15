package user_authentication

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type moduleMongoDataStore struct{
	db *mongo.Database
}

func (mmds moduleMongoDataStore) getUser(un string) user {
	var result user

	filter := bson.D{{"username", un}}

	err := mmds.db.Collection("Users").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
