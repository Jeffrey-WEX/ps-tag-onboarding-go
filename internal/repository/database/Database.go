package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName     = "user"
	databaseUsername = "root"
	databasePassword = "password"
)

type Database struct {
}

func NewDatabase() *mongo.Database {

	url := os.Getenv("DATABASE_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))

	if err != nil {
		panic(err)
	}

	return client.Database(databaseName)
}
