package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	url              = "mongodb://localhost:27017"
	databaseName     = "user"
	databaseUsername = "root"
	databasePassword = "password"
)

type Database struct {
}

func NewDatabase() *mongo.Database {

	// TODO: Add authentication
	// credential := options.Credential{
	// 	AuthMechanism: "SCRAM-SHA-1",
	// 	AuthSource:    "admin",
	// 	Username:      databaseUsername,
	// 	Password:      databasePassword,
	// }

	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url).SetAuth(credential))
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))

	if err != nil {
		panic(err)
	}

	return client.Database(databaseName)
}
