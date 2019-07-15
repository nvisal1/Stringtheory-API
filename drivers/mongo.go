package drivers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

var instance MongoConnection
var once sync.Once

type MongoConnection struct {
	Db *mongo.Database
	Client *mongo.Client
}

func (mc MongoConnection) Disconnect() {
	err := mc.Client.Disconnect(context.TODO())
	if err != nil {

	}
}

func Build() {
	once.Do(func() {
		db, client := connect()
		instance = MongoConnection{
			Db: db,
			Client: client,
		}
	})
}

func GetConnection() MongoConnection {
	if instance.Db == nil {
		log.Fatal("Database connection has not been established")
	}

	return instance
}

func connect() (*mongo.Database, *mongo.Client) {
	mongoURI, exists := os.LookupEnv("MONGO_URI")
	if exists {
		client, err := mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(mongoURI),
		)

		if err != nil {
			log.Fatal(err)
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}

		return client.Database("Cluster0"), client
	}
	return nil, nil
}