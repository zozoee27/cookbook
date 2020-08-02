package database

import (
	"context"
	"log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DatabaseClient *mongo.Client

func InitializeConnection () {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    var err error
    DatabaseClient, err = mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = DatabaseClient.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    log.Print("Connected to MongoDB on port 27017")
}

func Disconnect() {

    err := DatabaseClient.Disconnect(context.TODO())

    if err != nil {
        log.Fatal(err)
    }

    log.Print("Disconnected from MongoDB on port 27017")
}




