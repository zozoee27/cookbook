package database

import (
	"context"
	"log"

	"github.com/zozoee27/cookbook/backend/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoDb struct {
	userCollection *mongo.Collection
}

func CreateUserMongoDb(dbCollection *mongo.Collection) *UserMongoDb {
	return &UserMongoDb{
		userCollection: dbCollection}
}

func (d *UserMongoDb) Insert(user *entity.User) error {

	_, err := d.userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (d *UserMongoDb) FindOne(username string) (*entity.User, error) {
	result := &entity.User{}

	err := d.userCollection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
	if err != nil {
		log.Print("Find user error: ", err)
	}

	return result, err
}

func (d *UserMongoDb) ClearAllEntries() error {
	return d.userCollection.Drop(context.TODO())
}
