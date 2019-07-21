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

// Build is a wrapper for the connect function.
// This function ensures the integrity of the
// the singleton instance
func Build() {
	once.Do(func() {
		db, client := connect()
		instance = MongoConnection{
			Db: db,
			Client: client,
		}
	})
}

// GetConnection is an exported function
// that allows service modules to use
// the current instance of the database.
// If this function is called before Build,
// an error will be thrown
func GetConnection() MongoConnection {
	if instance.Db == nil {
		log.Fatal("Database connection has not been established")
	}

	return instance
}

// connect is responsible for establishing
// a connection with MongoDB
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