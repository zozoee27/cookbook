package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	DatabaseClient *mongo.Client
	Database       *mongo.Database
}

func InitializeConnection(dbName string) Connection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	var result Connection
	result.DatabaseClient, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = result.DatabaseClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	result.Database = result.DatabaseClient.Database(dbName)

	log.Print("Connected to MongoDB on port 27017")
	return result
}

func Disconnect(databaseClient *mongo.Client) {

	err := databaseClient.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Disconnected from MongoDB on port 27017")
}
