package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	url          = "mongodb://localhost:27017"
	databaseName = "user"
)

type Database struct {
}

func NewDatabase() *mongo.Database {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))

	if err != nil {
		panic(err)
	}

	return client.Database(databaseName)
}
