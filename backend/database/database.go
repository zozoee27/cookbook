package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Database       *mongo.Database
	DatabaseClient *mongo.Client
}

func (c *Connection) InitializeConnection(dbName string) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	c.DatabaseClient, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = c.DatabaseClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	c.Database = c.DatabaseClient.Database(dbName)

	log.Print("Connected to MongoDB on port 27017")
}

func (c *Connection) Disconnect() {

	err := c.DatabaseClient.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Disconnected from MongoDB on port 27017")
}
