package resource

import (
	"context"
	"fmt"

	"github.com/miscmo/note-mgr/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var NoteMgrMongo *mongo.Database

func InitMongo() {
	URI := "mongodb://admin:123456@172.17.0.4:27017"

	var (
		err    error
		client *mongo.Client
	)

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		panic(err)
	}

	NoteMgrMongo = client.Database("note_mgr")

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully connected to MongoDB!")
}

func GetDatabase() *mongo.Database {
	return NoteMgrMongo
}

func GetNoteCollection() *mongo.Collection {
	return GetDatabase().Collection((&entity.Note{}).CollName())
}

func GetElementCollection() *mongo.Collection {
	return GetDatabase().Collection((&entity.Element{}).CollName())
}
