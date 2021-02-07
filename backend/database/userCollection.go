package database

import (
	"context"
	"log"

	"github.com/zozoee27/cookbook/backend/account"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDatabase struct {
	UserCollection *mongo.Collection
}

func InitializeUserDatabase(dbCollection *mongo.Collection) *UserDatabase {
	return &UserDatabase{
		UserCollection: dbCollection}
}

func (d *UserDatabase) AddUserToCollection(user account.Account) error {

	_, err := d.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (d *UserDatabase) FindUserFromCollection(username string) account.Account {
	var result account.Account

	err := d.UserCollection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
	if err != nil {
		log.Print("Find user error: ", err)
	}

	return result
}

func (d *UserDatabase) ClearAllEntries() error {
	return d.UserCollection.Drop(context.TODO())
}
