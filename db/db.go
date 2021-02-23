package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// const dbName = "go-fiber"
// const mongoURI = "mongodb://localhost:27017/" + dbName

// Repo instnace
var Repo *MongoInstance

// Connect connect db
func Connect() error {
	var dbName = os.Getenv("DB_NAME")
	var mongoURI = os.Getenv("MONGODB_URI") + "/" + dbName
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db := client.Database(dbName)
	if err != nil {
		panic(err)
	}
	// defer client.Disconnect(ctx)
	Repo = &MongoInstance{
		Client: client,
		Db:     db,
	}
	return err
}
