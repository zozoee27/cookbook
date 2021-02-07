package database

import (
	"context"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/zozoee27/cookbook/backend/account"
)

var userDb *UserDatabase
var collection *mongo.Collection
var dbClient *mongo.Client

func TestMain(m *testing.M) {
	initDatabase()

	m.Run()

	userDb.ClearAllEntries()
	dbClient.Disconnect(context.TODO())
}

func TestAddValidUserToCollection(t *testing.T) {
	validAccount := account.Account{
		Username:  "ButtersButtons",
		Email:     "butters@buttons.com",
		FirstName: "Butters",
		LastName:  "Buttons",
		Password:  "Password123",
	}

	err := userDb.AddUserToCollection(validAccount)
	if err != nil {
		log.Fatal("FAILED - Could not add user to collection: ", err)
	}

	result := userDb.FindUserFromCollection("ButtersButtons")

	if result != validAccount {
		t.Errorf("Account information different")
	}
}

func initDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	dbClient, _ = mongo.Connect(context.TODO(), clientOptions)

	collection = dbClient.Database("CookbookDB_Test").Collection("users")
	userDb = InitializeUserDatabase(collection)
}
