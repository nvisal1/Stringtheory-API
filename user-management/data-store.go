package user_management

import "go.mongodb.org/mongo-driver/mongo"

type moduleMongoDataStore struct {
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
